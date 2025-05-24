package utilsAes

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateKey Function to generate a random 16-byte key as a base64-encoded string
func GenerateKey() string {
	key := make([]byte, 16)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(key)
}
