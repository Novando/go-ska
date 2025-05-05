package utilsJwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// CreateToken to issue JWT
func CreateToken(data map[string]interface{}, secret interface{}, opt Option) (string, error) {
	var (
		signingMethod = jwt.GetSigningMethod("HS256")
		duration      = 2 * time.Hour
	)

	// assign option if exist
	if opt.Method != nil {
		signingMethod = opt.Method
	}
	if opt.Duration != 0 {
		duration = opt.Duration
	}

	token := jwt.New(signingMethod)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	for key, value := range data {
		claims[key] = value
	}

	// Set expiration
	claims["exp"] = time.Now().Add(duration).Unix()
	claims["iat"] = time.Now().Unix()

	// Sign the token with the key
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
