package controller

import (
	"backend/utilService"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/machinebox/graphql"
)

type AuthResponse struct {
	ID              string `json:"id"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Email           string `json:"email"`
	Token           string `json:"token"`
	ProfileImageURL string `json:"profile_image_url"`
	Role            string `json:"role"`
}

type InputUser struct {
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type TokenInput struct {
	Token string `json:"token"`
}
type RequestBody struct {
	Action           map[string]interface{} `json:"action"`
	Input            InputUser              `json:"input"`
	RequestQuery     string                 `json:"request_query"`
	SessionVariables map[string]interface{} `json:"session_variables"`
}

type VerificationRequestBody struct {
	Action           map[string]interface{} `json:"action"`
	Input            TokenInput             `json:"input"`
	RequestQuery     string                 `json:"request_query"`
	SessionVariables map[string]interface{} `json:"session_variables"`
}


func sendToken(w http.ResponseWriter, id, firstName, lastName, email, profileImageURL, role string) {
	token, err := utilService.GetToken(id, "user")
	if err != nil {
		http.Error(w, "something went wrong when creating token", http.StatusBadRequest)
		return
	}

	response := AuthResponse{
		ID:              id,
		FirstName:       firstName,
		LastName:        lastName,
		Email:           email,
		Token:           token,
		ProfileImageURL: profileImageURL,
		Role:            role,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Send verification email
func sendVerificationEmail(email, token string) error {
	return utilService.SendEmail(email, token)
}

func Signup(w http.ResponseWriter, r *http.Request) {
	var reqBody RequestBody
	body, err := io.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(body, &reqBody); err != nil {
		respondWithError(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	newUser := reqBody.Input

	// Hash the password and use it later
	hashedPassword, err := utilService.HashPassword(newUser.Password)
	if err != nil {
		respondWithError(w, err.Error(), http.StatusBadRequest)
		return
	}

	client := graphql.NewClient("http://hasura:8080/v1/graphql")

	// Check if user already exists
	checkReq := graphql.NewRequest(`
        query($email: String!) {
            users(where: {email: {_eq: $email}}) {
                id
            }
        }
    `)
	checkReq.Var("email", newUser.Email)
	checkReq.Header.Set("x-hasura-admin-secret", "myadminsecretkey")

	var checkResp struct {
		Users []struct {
			ID string `json:"id"`
		} `json:"users"`
	}

	if err := client.Run(context.Background(), checkReq, &checkResp); err != nil {
		respondWithError(w, fmt.Sprintf("Error checking user existence: %v", err), http.StatusBadRequest)
		return
	}
	if len(checkResp.Users) > 0 {
		respondWithError(w, "User with this email already exists", http.StatusConflict)
		return
	}

	// Generate verification token
	verificationToken := uuid.New().String()

	// Store the user in the unverified_users table
	storeReq := graphql.NewRequest(`
        mutation($objects: [unverified_users_insert_input!]!) {
            insert_unverified_users(objects: $objects) {
                affected_rows
            }
        }
    `)

	storeReq.Var("objects", []map[string]interface{}{
		{
			"first_name":         newUser.FirstName,
			"last_name":          newUser.LastName,
			"email":              newUser.Email,
			"password":           hashedPassword,
			"verification_token": verificationToken,
		},
	})
	storeReq.Header.Set("x-hasura-admin-secret", "myadminsecretkey")

	var storeResp struct {
		InsertUnverifiedUsers struct {
			AffectedRows int `json:"affected_rows"`
		} `json:"insert_unverified_users"`
	}

	if err := client.Run(context.Background(), storeReq, &storeResp); err != nil {
		respondWithError(w, "Failed to store unverified user data", http.StatusInternalServerError)
		return
	}

	// Send verification email
	if err := sendVerificationEmail(newUser.Email, verificationToken); err != nil {
		respondWithError(w, "Failed to send verification email", http.StatusInternalServerError)
		return
	}

	// Send success response
	respondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Signup successful! Please check your email for verification.",
	})
}

// After the user verifies their email

func VerifyEmail(w http.ResponseWriter, r *http.Request) {
	var reqBody VerificationRequestBody
	body, err := io.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(body, &reqBody); err != nil {
		respondWithError(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	token := reqBody.Input.Token

	client := graphql.NewClient("http://hasura:8080/v1/graphql")

	// Retrieve the user from the unverified_users table
	checkReq := graphql.NewRequest(`
        query($token: String!) {
            unverified_users(where: {verification_token: {_eq: $token}}) {
                first_name
                last_name
                email
                password
            }
        }
    `)
	checkReq.Var("token", token)
	checkReq.Header.Set("x-hasura-admin-secret", "myadminsecretkey")

	var checkResp struct {
		UnverifiedUsers []struct {
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Email     string `json:"email"`
			Password  string `json:"password"`
		} `json:"unverified_users"`
	}

	if err := client.Run(context.Background(), checkReq, &checkResp); err != nil {
		respondWithError(w, "Error retrieving unverified user", http.StatusBadRequest)
		return
	}

	if len(checkResp.UnverifiedUsers) == 0 {
		respondWithError(w, "Invalid or expired token", http.StatusBadRequest)
		return
	}

	userData := checkResp.UnverifiedUsers[0]

	// Insert the user into the main users table
	insertReq := graphql.NewRequest(`
        mutation($objects: [users_insert_input!]!) {
            insert_users(objects: $objects) {
                returning {
                    id
                    first_name
                    last_name
                    email
                    role
                    profile_image_url
                }
            }
        }
    `)

	insertReq.Var("objects", []map[string]interface{}{
		{
			"first_name": userData.FirstName,
			"last_name":  userData.LastName,
			"email":      userData.Email,
			"password":   userData.Password,
		},
	})
	insertReq.Header.Set("x-hasura-admin-secret", "myadminsecretkey")

	var insertResp struct {
		InsertUsers struct {
			Returning []struct {
				ID              string `json:"id"`
				FirstName       string `json:"first_name"`
				LastName        string `json:"last_name"`
				Email           string `json:"email"`
				ProfileImageURL string `json:"profile_image_url"`
				Role            string `json:"role"`
			} `json:"returning"`
		} `json:"insert_users"`
	}

	if err := client.Run(context.Background(), insertReq, &insertResp); err != nil {
		http.Error(w, fmt.Sprintf("Error moving user to main table: %v", err), http.StatusBadRequest)
		return
	}

	// Delete the user from the unverified_users table
	deleteReq := graphql.NewRequest(`
        mutation($token: String!) {
            delete_unverified_users(where: {verification_token: {_eq: $token}}) {
                affected_rows
            }
        }
    `)
	deleteReq.Var("token", token)
	deleteReq.Header.Set("x-hasura-admin-secret", "myadminsecretkey")

	if err := client.Run(context.Background(), deleteReq, nil); err != nil {
		http.Error(w, fmt.Sprintf("Error deleting unverified user: %v", err), http.StatusInternalServerError)
		return
	}

	verifiedUser := insertResp.InsertUsers.Returning[0]

	// Send authentication token after email verification
	sendToken(w, verifiedUser.ID, verifiedUser.FirstName, verifiedUser.LastName, verifiedUser.Email, verifiedUser.ProfileImageURL, verifiedUser.Role)

	w.WriteHeader(http.StatusOK)
}


func Login(w http.ResponseWriter, r *http.Request) {
	var reqBody RequestBody
	// Decode the request body
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, `{"error": "Invalid JSON format"}`, http.StatusBadRequest)
		return
	}

	newUser := reqBody.Input

	// Prepare the GraphQL query
	client := graphql.NewClient("http://hasura:8080/v1/graphql")
	req := graphql.NewRequest(`
        query($email: String!) {
            users(where: {email: {_eq: $email}}) {
                id
                first_name
				last_name
                email
                password
                profile_image_url
				role
            }
        }
    `)

	req.Var("email", newUser.Email)
	req.Header.Set("x-hasura-admin-secret", "myadminsecretkey")
	// Execute the GraphQL query
	var respData struct {
		Users []struct {
			ID              string `json:"id"`
			FirstName       string `json:"first_name"`
			LastName        string `json:"last_name"`
			Email           string `json:"email"`
			Password        string `json:"password"`
			Role            string `json:"role"`
			ProfileImageURL string `json:"profile_image_url"` // Include profile image URL in the response
		} `json:"users"`
	}

	if err := client.Run(context.Background(), req, &respData); err != nil {
		http.Error(w, `{"error": "Error querying user data"}`, http.StatusInternalServerError)
		fmt.Println("Error querying user data:", err)
		return
	}

	// Check if the user exists and validate password
	if len(respData.Users) == 0 {
		http.Error(w, `{"error": "Invalid credentials"}`, http.StatusBadRequest)
		return
	}

	user := respData.Users[0]
	if !utilService.ComparePasswords(user.Password, newUser.Password) {
		http.Error(w, `{"error": "Invalid credentials"}`, http.StatusBadRequest)
		return
	}

	// Send authentication token
	sendToken(w, user.ID, user.FirstName, user.LastName, user.Email, user.ProfileImageURL, user.Role)
}

func respondWithError(w http.ResponseWriter, message string, statusCode int) {
	respondWithJSON(w, statusCode, map[string]string{"message": message})
}

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}
