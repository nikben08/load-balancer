package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func GenerateHash(password string) (string, error) {
	passwordSalt := []byte(password)
	passwordHash, hashError := bcrypt.GenerateFromPassword(passwordSalt, bcrypt.DefaultCost)

	if hashError != nil {
		return "", hashError // Return an empty string and the error
	}

	return string(passwordHash), nil
}

func CheckPassword(passwordHash []byte, password string) error {
	passwordSalt := []byte(password)

	compareError := bcrypt.CompareHashAndPassword(passwordHash, passwordSalt)
	if compareError != nil {
		return compareError
	}

	return nil
}
