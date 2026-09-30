package helpers

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

var (
	PASSWORD_SALT = os.Getenv("PASSWORD_SALT")
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(fmt.Sprintf("%s;%s", PASSWORD_SALT, password)), 12)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(fmt.Sprintf("%s;%s", PASSWORD_SALT, password)))
	return err == nil
}

// MinPasswordLength is the shared policy for every account-creation path.
const MinPasswordLength = 12

// ValidatePassword checks length and bcrypt's 72-byte input limit (salt + ";" + password).
func ValidatePassword(password string) error {
	if len(password) < MinPasswordLength {
		return fmt.Errorf("password must contain at least %d characters", MinPasswordLength)
	}
	if len([]byte(PASSWORD_SALT+";"+password)) > 72 {
		return fmt.Errorf("password is too long for bcrypt's 72-byte input limit")
	}
	return nil
}
