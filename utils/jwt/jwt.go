package utils

import (
	"fmt"
	"load-balancer/config"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func GenerateJwtToken(userID uuid.UUID, Email string, Name string) (string, error) {
	var secretKey = []byte(config.GetJwtSecretKey())
	var expirationMinutes, err = strconv.Atoi(config.GetJwtExpirationMinutes())

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["email"] = Email
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(expirationMinutes)).Unix()
	claims["userId"] = userID

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println(err)
	}
	return tokenString, err
}
