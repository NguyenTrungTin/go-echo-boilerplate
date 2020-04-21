package auth

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/nguyentrungtin/go-echo-boilerplate/config"
)

func NewToken(id int, username string, role string) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Get Expire time
	exp, _ := strconv.Atoi(config.JWT_EXP)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(int64(exp))).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.JWT_KEY))
	if err != nil {
		return "", err
	}

	signedToken := "Bearer " + t

	return signedToken, nil
}

func DeveloperToken(id int, username string, role string) (string, error) {

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["iat"] = time.Now().Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.JWT_KEY))
	if err != nil {
		return "", err
	}

	signedToken := "Bearer " + t

	return signedToken, nil
}

func MuvitToken(authToken, authSecret string) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["authToken"] = authToken
	claims["iat"] = time.Now().Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(authSecret))
	if err != nil {
		return "", err
	}

	signedToken := "Bearer " + t

	return signedToken, nil
}
