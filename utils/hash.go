package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string,error) {

	hPass, err := bcrypt.GenerateFromPassword([]byte(password),14)

	if(err != nil){
		return "",err
	}

	return string(hPass),nil
}

func CheckHash(password,Haspass string) bool{

	err := bcrypt.CompareHashAndPassword([]byte(Haspass),[]byte(password))

	return err == nil
}