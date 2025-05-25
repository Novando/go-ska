package utilsJwt

import (
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
	IssuedAt       int64       `json:"iat"`
	ExpirationTime int64       `json:"exp"`
	Data           IDTokenData `json:"data"`
}

type Option struct {
	Method   jwt.SigningMethod
	Duration time.Duration
}
