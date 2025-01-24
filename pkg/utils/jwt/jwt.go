package utilsJwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type IDTokenData struct {
	Email      string `json:"email"`
	FullName   string `json:"full_name"`
	LastName   string `json:"last_name"`
	FirstName  string `json:"first_name"`
	PictureURL string `json:"picture_url"`
}

type IDToken struct {
	Issuer         string      `json:"iss"`
	Subject        string      `json:"sub"`
	Audience       string      `json:"aud"`
	IssuedAt       time.Time   `json:"iat"`
	ExpirationTime time.Time   `json:"exp"`
	Data           IDTokenData `json:"data"`
}

// CreateToken to issue JWT
func CreateToken(data map[string]interface{}, secret string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	for key, value := range data {
		claims[key] = value
	}

	// Set expiration
	claims["exp"] = time.Now().Add(2 * time.Hour)
	claims["iat"] = time.Now().Unix()

	// Sign the token with the key
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseToken to validate and decrypt the data
func ParseToken(tokenString, secret string) (map[string]interface{}, error) {
	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	// Verify token
	if !token.Valid {
		return nil, fmt.Errorf("token is invalid")
	}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("failed to parse claims")
	}

	return claims, nil
}

// GetValue validate and get single value from JWT
func GetValue(token, key, secret string) (res interface{}, err error) {
	claims, err := ParseToken(token, secret)
	if err != nil {
		return
	}
	res = claims[key]
	if res == nil {
		err = fmt.Errorf("key not found")
	}
	return
}