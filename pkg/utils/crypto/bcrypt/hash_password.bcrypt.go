package utilsBcrypt

import (
	"github.com/novando/go-ska/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword Function to hash a password
func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.Call().Errorf(err.Error())
	}
	return string(hashedPassword)
}
