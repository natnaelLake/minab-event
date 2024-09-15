package utilService

import "golang.org/x/crypto/bcrypt"

func ComparePasswords(hashedPassword string, password string) bool {
	// Convert the hashed password string to a byte slice
	hashedPasswordBytes := []byte(hashedPassword)
	// Compare the hashed password byte slice with the plaintext password byte slice
	err := bcrypt.CompareHashAndPassword(hashedPasswordBytes, []byte(password))
	return err == nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
