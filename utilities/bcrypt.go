package utilities

import (
	"golang.org/x/crypto/bcrypt"
)

// Create string with hash
func CompareHashes(originalValue, hash string) bool {
	hashError := bcrypt.CompareHashAndPassword([]byte(hash), []byte(originalValue))
	return hashError == nil
}

// Create hash from a string
func MakeHash(value string) (string, error) {
	bytes, hashError := bcrypt.GenerateFromPassword([]byte(value), 10)
	return string(bytes), hashError
}
