package utils

import (
	"crypto/rand"
	"encoding/base64"
)

// RememberTokenBytes is the number of bytes that will be generated.
const RememberTokenBytes = 32

// GenerateBytes function is used to generated random n bytes, or will
// return an error if one, this uses crypto/rand
func GenerateBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GenerateString function is used to generate a byte slice of size nBytes and
// then it will return a string that's the base64 url encoded version.
// of that byte slice.
func GenerateString(nBytes int) (string, error) {
	b, err := GenerateBytes(nBytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// RememberToken function is used to generate a string using the remember token bytes.
func RememberToken() (string, error) {
	return GenerateString(RememberTokenBytes)
}
