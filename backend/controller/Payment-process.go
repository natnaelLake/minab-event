package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Define the input structure for payment processing
type ProcessPaymentInput struct {
	TxRef      string  `json:"tx_ref"`
	TotalPrice float64 `json:"total_price"`
	EventID    string  `json:"event_id"`
	UserID     string  `json:"user_id"`
	Email      string  `json:"email"`
	FirstName  string  `json:"first_name"`
	LastName   string  `json:"last_name"`
	Quantity   int     `json:"quantity"`
	ReturnUrl  string  `json:"return_url"`
}

// Define the full request structure received from Hasura Actions
type ProcessPaymentRequest struct {
	Action map[string]interface{} `json:"action"`
	Input  struct {
		UserInput ProcessPaymentInput `json:"input"`
	} `json:"input"`
	RequestQuery     string                 `json:"request_query"`
	SessionVariables map[string]interface{} `json:"session_variables"`
}

// Define the structure of the response from the payment API
type ProcessPaymentResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Data    struct {
		CheckoutURL string `json:"checkout_url"`
	} `json:"data"`
}

// ProcessPaymentHandler handles the payment process
func ProcessPaymentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error": "Invalid request method"}`, http.StatusMethodNotAllowed)
		return
	}

	var reqBody ProcessPaymentRequest
	// Decode the incoming JSON request body
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, `{"error": "Failed to decode request"}`, http.StatusBadRequest)
		return
	}

	paymentInput := reqBody.Input.UserInput

	url := "https://api.chapa.co/v1/transaction/initialize"
	method := "POST"

	// Create the payload for the payment API using the input parameters
	payload := map[string]interface{}{
		"amount":       paymentInput.TotalPrice * float64(paymentInput.Quantity), 
		"email":        paymentInput.Email,     
		"first_name":   paymentInput.FirstName, 
		"last_name":    paymentInput.LastName,  
		"phone_number": "0912345678",          
		"tx_ref":       paymentInput.TxRef,
		"callback_url": "https://webhook.site/077164d6-29cb-40df-ba29-8a00e59a7e60",
		"return_url":   paymentInput.ReturnUrl,
		"customization": map[string]string{
			"title":       "Pay Order",       
			"description": "I love payments", 
		},
		"meta": map[string]string{
			"hide_receipt": "true",
		},
	}

	// Convert the payload to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, `{"error": "Failed to encode request payload"}`, http.StatusInternalServerError)
		return
	}
	payloadReader := bytes.NewReader(payloadBytes)

	client := &http.Client{}

	// Create a new HTTP request with the payload
	req, err := http.NewRequest(method, url, payloadReader)
	if err != nil {
		http.Error(w, `{"error": "Failed to create request"}`, http.StatusInternalServerError)
		return
	}
	req.Header.Add("Authorization", "Bearer CHASECK_TEST-m4enUTE37kbMMK9f3L3tD4WpfDrfGhKh") // Use your valid Chapa API key
	req.Header.Add("Content-Type", "application/json")

	// Send the request
	res, err := client.Do(req)
	if err != nil {
		http.Error(w, `{"error": "Failed to send request"}`, http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	// Check for non-200 response status
	if res.StatusCode != http.StatusOK {
		responseBody, _ := io.ReadAll(res.Body)
		fmt.Printf("Received non-200 response: %s\n", string(responseBody))
		http.Error(w, fmt.Sprintf(`{"error": "Received non-200 response: %d"}`, res.StatusCode), http.StatusInternalServerError)
		return
	}

	// Read and decode the response body from the payment API
	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		http.Error(w, `{"error": "Failed to read response body"}`, http.StatusInternalServerError)
		return
	}

	// Log the response body for debugging
	fmt.Printf("Response body: %s\n", string(responseBody))

	var response ProcessPaymentResponse
	if err := json.Unmarshal(responseBody, &response); err != nil {
		http.Error(w, `{"error": "Failed to decode response"}`, http.StatusInternalServerError)
		return
	}

	// Send the response back to the client
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, `{"error": "Failed to encode response"}`, http.StatusInternalServerError)
	}
}
