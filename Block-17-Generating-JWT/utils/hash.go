package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {

	// The value of 14 controls how complex the hashing process should be.
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func IsPasswordValid(normalPassword, retrievedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(normalPassword), []byte(retrievedPassword))
	
	return err == nil
}