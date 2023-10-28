package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashPassword(password []byte) string {
	hashed, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)

	if err != nil {
		log.Fatalf("Error while hashing password: %v", err)
		return ""
	}

	return string(hashed)
}
