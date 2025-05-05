package utilsAes

import (
	"bytes"
	"crypto/aes"
)

// PKCS7Pad adds padding to the plaintext
func PKCS7Pad(plainText []byte) []byte {
	padding := aes.BlockSize - (len(plainText) % aes.BlockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(plainText, padText...)
}
