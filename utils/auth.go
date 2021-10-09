package utils

import "golang.org/x/crypto/bcrypt"

func HashPass(pwd string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(pwd), 12)
	return string(hash)
}

func CompHash(hashpwd, plainpwd string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashpwd), []byte(plainpwd)); err != nil {
		return false
	}

	return true
}