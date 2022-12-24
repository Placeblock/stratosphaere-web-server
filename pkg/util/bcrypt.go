package util

import "golang.org/x/crypto/bcrypt"

func CompareHash(hashed string, plain []byte) bool {
	byteHash := []byte(hashed)
	err := bcrypt.CompareHashAndPassword(byteHash, plain)
	return err == nil
}

func GenerateHash(password []byte) []byte {
	hash, _ := bcrypt.GenerateFromPassword(password, 12)
	return hash
}
