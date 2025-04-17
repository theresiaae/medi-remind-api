package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword digunakan untuk mengenkripsi password sebelum disimpan
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash digunakan untuk membandingkan password yang diinput dengan hash di database
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
