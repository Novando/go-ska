package utilsBcrypt

import "golang.org/x/crypto/bcrypt"

// ValidatePassword Function to validate a password against its hash
func ValidatePassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
