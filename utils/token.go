package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"os"
	"time"
)

func GenerateToken(email string) string {
	mySigningKey := []byte(os.Getenv("TOKEN_SECRET_KEY"))
	expiredTime, _ := time.ParseDuration(os.Getenv("EXPIRED_TIME"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(expiredTime).Unix(),
	})
	tokenKey, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Println("Generate token key error")
		log.Println(err)
	}
	return tokenKey
}

func ValidateToken(tokenString string) (bool, string) {
	mySigningKey := []byte(os.Getenv("TOKEN_SECRET_KEY"))
	var email string

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		log.Println("Validate token key error")
		log.Println(err)
		return false, email
	}

	for key, value := range claims {
		if key == "email" {
			email = fmt.Sprintf("%v",value)
		}
	}
	return token.Valid, email
}
