package security

import (
	"crypto/rand"
	"encoding/hex"
	"strings"
)

func GenerateRandomString(length int) string {
	b := make([]byte, length/2)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(b)
}

func TrimString(s string) string {
	return strings.TrimSpace(s)
}
