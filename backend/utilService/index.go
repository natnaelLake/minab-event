package utilService

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/smtp"
	"regexp"
	"sync"
	"time"
)

var (
	tokenStore = make(map[string]string)
	mu         sync.Mutex
)

// GenerateVerificationToken generates a unique, cryptographically secure token for email verification.
func GenerateVerificationToken(email string) (string, error) {
	// Generate 32 random bytes
	tokenBytes := make([]byte, 32)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	// Encode as a base64 string
	token := base64.URLEncoding.EncodeToString(tokenBytes)

	// Optionally, store token with email for verification later
	StoreVerificationToken(email, token)

	return token, nil
}

// SendVerificationEmail sends a verification email to the user's email with the provided token.
func SendVerificationEmail(email, token string) error {
	// Set up SMTP server (for example, using Gmail's SMTP)
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	senderEmail := "your-email@gmail.com"
	password := "your-password"

	// Message body
	message := []byte(fmt.Sprintf("To verify your email, please use the following token or click the link: %s", token))

	// Authentication
	auth := smtp.PlainAuth("", senderEmail, password, smtpHost)

	// Send email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, senderEmail, []string{email}, message)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

// StoreVerificationToken stores the email and its corresponding token temporarily.
func StoreVerificationToken(email, token string) {
	mu.Lock()
	defer mu.Unlock()

	// Store the token with an expiration time (in this example, tokens are valid for 15 minutes)
	tokenStore[email] = token

	// Clean up the token after 15 minutes
	go func() {
		time.Sleep(15 * time.Minute)
		mu.Lock()
		delete(tokenStore, email)
		mu.Unlock()
	}()
}

// VerifyToken checks if the provided token matches the stored token for the given email.
func VerifyToken(email, token string) bool {
	mu.Lock()
	defer mu.Unlock()

	storedToken, exists := tokenStore[email]
	if !exists {
		return false
	}

	// Check if the token matches
	return storedToken == token
}

// IsValidEmail checks if the email is in a valid format using regex.
func IsValidEmail(email string) bool {
	// Basic email validation regex
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}