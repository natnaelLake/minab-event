package utilService

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func SendEmail(email string, token string) error {
	apiURL := "https://sandbox.api.mailtrap.io/api/send/3139740"
	apiToken := "1d2778a04f0eb13bcc35852ae487c418"

	if apiURL == "" || apiToken == "" {
		return fmt.Errorf("missing email API configuration")
	}

	emailData, err := prepareEmailData(email, token)
	if err != nil {
		fmt.Printf("Error preparing email data: %v\n", err)
		return fmt.Errorf("failed to prepare email data: %w", err)
	}

	err = sendEmailRequest(apiURL, apiToken, emailData)
	if err != nil {
		fmt.Printf("Error sending email request: %v\n", err)
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

func sendEmailRequest(url, apiToken string, emailData map[string]interface{}) error {
	jsonData, err := json.Marshal(emailData)
	if err != nil {
		fmt.Printf("Error marshaling email data: %v\n", err)
		return fmt.Errorf("failed to marshal email data: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Debug: Check the status code and response body
	fmt.Printf("Response Status Code: %d\n", resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("Error Response Body: %s\n", string(body))
		return fmt.Errorf("failed to send email, status code: %d, response: %s", resp.StatusCode, body)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// Debug: Print response body
	fmt.Printf("Response Body: %s\n", string(body))

	return nil
}

func prepareEmailData(email, token string) (map[string]interface{}, error) {
	var subject, htmlContent string

	subject = "Confirm Your Email"
	htmlContent = generateConfirmationEmailHTML(token)

	emailData := map[string]interface{}{
		"from": map[string]string{
			"email": "mailtrap@example.com",
			"name":  "Mailtrap Test",
		},
		"to": []map[string]string{
			{
				"email": email,
			},
		},
		"subject": subject,
		"html":    htmlContent,
	}
	return emailData, nil
}

func generateConfirmationEmailHTML(token string) string {
	confirmationLink := fmt.Sprintf("http://localhost:3000/verify-email?token=%s", token)
	return fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8" />
			<meta name="viewport" content="width=device-width, initial-scale=1.0" />
			<title>Email Confirmation</title>
			<style>
				/* Styles omitted for brevity */
			</style>
		</head>
		<body>
			<div class="container">
				<div class="header">
					<img src="https://via.placeholder.com/100x100" alt="Company Logo" />
					<h1>Email Confirmation</h1>
				</div>
				<div class="content">
					<p>Thank you for signing up! Please confirm your email address by clicking the button below:</p>
					<a href="%s" class="button">Confirm Email</a>
				</div>
			</div>
		</body>
		</html>
	`, confirmationLink)
}
