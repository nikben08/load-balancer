<<<<<<< HEAD
package utils

import (
	"strconv"
	"time"
	"load-balancer/config"

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

=======
func generateJWT(username string, userId int, accessLevel int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":      userId,
		"username":    username,
		"accessLevel": accessLevel,
	})
	secretKey := []byte(config.Config("JWT_SECRET_FOR_LOCAL"))
>>>>>>> 3e42eaab96636658b0510040657ebdb968aa0125
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println(err)
	}
	return tokenString, err
}