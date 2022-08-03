package plugins

import (
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	cost := os.Getenv("ENCRYPT_COST")
	costInt, _ := strconv.Atoi(cost)
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), costInt)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
