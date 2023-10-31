package handlers

import (
	"fmt"
	"load-balancer/contracts"
	"load-balancer/models"
	"load-balancer/repositories"
	"load-balancer/services"
	encryption "load-balancer/utils/encryption"
	jwt "load-balancer/utils/jwt"

	"github.com/gofiber/fiber/v2"
)

type User models.User

func Login(c *fiber.Ctx) error {
	json := new(contracts.LoginRequest)
	if err := c.BodyParser(json); err != nil {
		fmt.Println(err)
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
			"error":   err.Error(),
		})
	}

	found := &models.User{Email: json.Email}
	err := repositories.GetUserByEmail(found)
	if err != nil {
		return c.JSON("Wrong username or password")
	}

	if encryption.ComparePasswords(found.Hash, []byte(json.Password)) {
		token, _ := jwt.GenerateJwtToken(found.Id, found.Email, found.Name)
		var response = contracts.AuthResponse{
			Code:    200,
			Message: "User successfully logged",
			Token:   token,
		}
		return c.JSON(response)
	} else {
		return c.JSON("Wrong username or password")
	}
}

func Signup(c *fiber.Ctx) error {
	json := new(contracts.SignupRequest)
	if err := c.BodyParser(json); err != nil {
		fmt.Println(err)
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
			"error":   err.Error(),
		})
	}

	if json.Password != json.PasswordRepeat {
		return c.JSON("passwords do not match")
	}
	fmt.Println(json)
	hash, _ := encryption.GenerateHash([]byte(json.Password))
	var newUser = &models.User{
		Email:   json.Email,
		Hash:    hash,
		Name:    json.Name,
		Surname: json.Surname,
	}

	err := repositories.CreateNewUser(newUser)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
			"error":   err.Error(),
		})
	}

	token, _ := jwt.GenerateJwtToken(newUser.Id, json.Email, json.Name)

	var response = contracts.AuthResponse{
		Code:    200,
		Message: "User successfully created",
		Token:   token,
	}

	return c.JSON(response)
}

func GoogleOAuth(c *fiber.Ctx) error {
	json := new(contracts.GoogleOAuthRequest)
	if err := c.BodyParser(json); err != nil {
		fmt.Println(err)
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
			"error":   err.Error(),
		})
	}

	token := services.HandleGoogleOAuth(*json)

	var response = contracts.AuthResponse{
		Code:    200,
		Message: "User successfully created",
		Token:   token,
	}

	return c.JSON(response)
}
