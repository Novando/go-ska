package utilsAes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

// Encrypt Function to perform AES encryption
func Encrypt(plainText string, key string) string {
	keyBytes, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		panic(err)
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		panic(err)
	}

	// Convert plaintext string to byte slice
	plainTextBytes := []byte(plainText)

	// Pad the plaintext
	plainTextBytes = PKCS7Pad(plainTextBytes)

	cipherText := make([]byte, aes.BlockSize+len(plainTextBytes))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[aes.BlockSize:], plainTextBytes)

	// Encode ciphertext to base64 string
	return base64.StdEncoding.EncodeToString(cipherText)
}
