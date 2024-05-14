package helpers

import (
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	salt := os.Getenv("BCRYPT_SALT")
	saltInt, _ := strconv.Atoi(salt)

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), saltInt)
	return string(bytes), err
}

func ComparePassword(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
