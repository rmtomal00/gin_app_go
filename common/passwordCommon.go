package common

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pass string)(string, error){
	if pass == "" || len(pass) < 6 {
		return  "", fmt.Errorf("%s", "password can't be null");
	}
	hashPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err;
	}
	return string(hashPass), nil
}

func IsValidPass(hash string, pass string) bool{
	if hash == "" || pass == "" {
		return false
	}
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass));

	if err != nil {
		return false
	}
	return true
}