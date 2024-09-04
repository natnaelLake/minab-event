package controller

import (
	"backend/utilService"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/machinebox/graphql" // Import the GraphQL client package
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

type RequestBody struct {
	Action           map[string]interface{} `json:"action"`
	Input            InputUser              `json:"input"` // Corrected to match Hasura's payload structure
	RequestQuery     string                 `json:"request_query"`
	SessionVariables map[string]interface{} `json:"session_variables"`
}

func sendToken(w http.ResponseWriter, id string, first_name string, last_name string, email string, profileImageURL string, role string) {
	token, err := utilService.GetToken(id, "user")
	if err != nil {
		http.Error(w, "something went wrong when creating token", http.StatusBadRequest)
		return
	}

	response := AuthResponse{
		ID:              id,
		FirstName:       first_name,
		LastName:        last_name,
		Email:           email,
		Token:           token,
		ProfileImageURL: profileImageURL,
		Role:            role,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func Signup(w http.ResponseWriter, r *http.Request) {
	var reqBody RequestBody
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	r.Body = io.NopCloser(bytes.NewBuffer(body))

	if err := json.Unmarshal(body, &reqBody); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	newUser := reqBody.Input
	fmt.Print("++++++++++++++++++++++++", newUser)
	hashedPassword, err := utilService.HashPassword(newUser.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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

	var checkResp struct {
		Users []struct {
			ID string `json:"id"`
		} `json:"users"`
	}

	if err := client.Run(context.Background(), checkReq, &checkResp); err != nil {
		http.Error(w, fmt.Sprintf("Error checking user existence: %v", err), http.StatusBadRequest)
		return
	}

	if len(checkResp.Users) > 0 {
		http.Error(w, "User with this email already exists", http.StatusConflict)
		return
	}

	// Proceed with user creation
	req := graphql.NewRequest(`
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

	req.Var("objects", []map[string]interface{}{
		{
			"first_name": newUser.FirstName,
			"last_name":  newUser.LastName,
			"email":      newUser.Email,
			"password":   hashedPassword,
		},
	})

	var respData struct {
		InsertUsers struct {
			Returning []struct {
				ID              string `json:"id"`
				FirstName       string `json:"first_name"`
				LastName        string `json:"last_name"`
				Email           string `json:"email"`
				Role            string `json:"role"`
				ProfileImageURL string `json:"profile_image_url"`
			} `json:"returning"`
		} `json:"insert_users"`
	}

	if err := client.Run(context.Background(), req, &respData); err != nil {
		http.Error(w, fmt.Sprintf("Error executing GraphQL mutation: %v", err), http.StatusBadRequest)
		return
	}

	if len(respData.InsertUsers.Returning) == 0 {
		http.Error(w, "No user ID returned", http.StatusBadRequest)
		return
	}

	user := respData.InsertUsers.Returning[0]
	sendToken(w, user.ID, user.FirstName, user.LastName, user.Email, user.ProfileImageURL, user.Role)
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

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var inputUser struct {
		Email       string `json:"email"`
		Password    string `json:"password"`
		NewPassword string `json:"newPassword"`
	}

	if err := json.NewDecoder(r.Body).Decode(&inputUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Prepare the GraphQL query to find the user
	client := graphql.NewClient("http://hasura:8080/v1/graphql")
	req := graphql.NewRequest(`
        query($email: String!) {
            users(where: {email: {_eq: $email}}) {
                id
                first_name
				last_name
				role
                email
                password
                profile_image_url
            }
        }
    `)

	req.Var("email", inputUser.Email)

	// Execute the GraphQL query
	var queryData struct {
		Users []struct {
			ID              string `json:"id"`
			FirstName       string `json:"first_name"`
			LastName        string `json:"last_name"`
			Email           string `json:"email"`
			Password        string `json:"password"`
			ProfileImageURL string `json:"profile_image_url"` // Include profile image URL in the response
		} `json:"users"`
	}

	if err := client.Run(context.Background(), req, &queryData); err != nil {
		http.Error(w, "Error querying user data", http.StatusInternalServerError)
		fmt.Println("Error querying user data:", err)
		return
	}

	// Validate user credentials
	if len(queryData.Users) == 0 || !utilService.ComparePasswords(queryData.Users[0].Password, inputUser.Password) {
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
		return
	}

	// Hash the new password if provided
	var newPassword string
	if inputUser.NewPassword != "" {
		var err error
		newPassword, err = utilService.HashPassword(inputUser.NewPassword)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		newPassword = queryData.Users[0].Password
	}

	// Prepare the GraphQL mutation to update the user
	mutation := graphql.NewRequest(`
        mutation($email: String!, $password: String!) {
            update_users(
                where: {email: {_eq: $email}},
                _set: {password: $password}
            ) {
                returning {
                    id
                    first_name,
					last_name
                    email
                    profile_image_url
					role
                }
            }
        }
    `)

	mutation.Var("email", inputUser.Email)
	mutation.Var("password", newPassword)

	// Execute the GraphQL mutation
	var mutationData struct {
		UpdateUsers struct {
			Returning []struct {
				ID              string `json:"id"`
				FirstName       string `json:"first_name"`
				LastName        string `json:"last_name"`
				Email           string `json:"email"`
				Role            string `json:"role"`
				ProfileImageURL string `json:"profile_image_url"` // Include profile image URL in the response
			} `json:"returning"`
		} `json:"update_users"`
	}

	if err := client.Run(context.Background(), mutation, &mutationData); err != nil {
		http.Error(w, "Error updating user data", http.StatusInternalServerError)
		fmt.Println("Error updating user data:", err)
		return
	}

	// Check if the user ID was returned
	if len(mutationData.UpdateUsers.Returning) == 0 {
		http.Error(w, "No user ID returned", http.StatusBadRequest)
		fmt.Println("Error: No user ID returned")
		return
	}

	// Send updated user information
	user := mutationData.UpdateUsers.Returning[0]
	sendToken(w, user.ID, user.FirstName, user.LastName, user.Email, user.ProfileImageURL, user.Role)
}
