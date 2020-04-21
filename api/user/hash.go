package user

import (
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword is used to hash user password using bcrypt
func HashPassword(pw string) (string, error) {
	pwBytes := []byte(pw)
	hashBytes, err := bcrypt.GenerateFromPassword(pwBytes, bcrypt.DefaultCost)
	if err != nil {
		log.Error("ERROR: Error happen when hashing password")
		log.Error(err)
		return "", err
	}
	return string(hashBytes), nil
}
