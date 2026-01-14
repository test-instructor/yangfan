package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func SecureRandomString(byteLen int) (string, error) {
	if byteLen <= 0 {
		byteLen = 16
	}
	b := make([]byte, byteLen)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}
