package utilsJwt

import "fmt"

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
