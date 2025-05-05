package utilsAes

// PKCS7Unpad removes padding from the plaintext
func PKCS7Unpad(plainText []byte) []byte {
	padding := int(plainText[len(plainText)-1])
	return plainText[:len(plainText)-padding]
}
