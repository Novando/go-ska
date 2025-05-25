package utilsAes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// Decrypt Function to perform AES decryption
func Decrypt(cipherText string, key string) string {
	keyBytes, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		panic(err)
	}

	cipherTextBytes, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		panic(err)
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		panic(err)
	}

	if len(cipherTextBytes) < aes.BlockSize {
		panic("cipherText too short")
	}

	iv := cipherTextBytes[:aes.BlockSize]
	cipherTextBytes = cipherTextBytes[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherTextBytes, cipherTextBytes)

	// Remove PKCS7 padding
	plainTextBytes := PKCS7Unpad(cipherTextBytes)

	// Convert byte slice to string
	return string(plainTextBytes)
}
