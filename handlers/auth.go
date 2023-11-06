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
		return c.Status(400).JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
			"error":   err.Error(),
		})
	}

	found := &models.User{Email: json.Email}
	err := repositories.GetUserByEmail(found)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"code":    401,
			"message": "Wrong email or password",
			"error":   err.Error(),
		})
	}

	if encryption.ComparePasswords(found.Hash, []byte(json.Password)) {
		token, _ := jwt.GenerateJwtToken(found.Id, found.Email, found.Name)
		return c.Status(200).JSON(fiber.Map{
			"code":         200,
			"message":      "User successfully logged",
			"access_token": token,
		})
	} else {
		return c.Status(401).JSON(fiber.Map{
			"code":    401,
			"message": "Wrong email or password",
			"error":   err.Error(),
		})
	}
}

func Signup(c *fiber.Ctx) error {
	json := new(contracts.SignupRequest)
	if err := c.BodyParser(json); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
			"error":   err.Error(),
		})
	}

	if json.Password != json.PasswordRepeat {
		return c.Status(401).JSON(fiber.Map{
			"code":    401,
			"message": "Passwords do not match",
			"error":   "INVALID_PASSWORD",
		})
	}

	hash, _ := encryption.GenerateHash([]byte(json.Password))

	var newUser = &models.User{
		Email:   json.Email,
		Hash:    hash,
		Name:    json.Name,
		Surname: json.Surname,
	}

	err := repositories.CreateNewUser(newUser)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code":    400,
			"message": "Something went wrong",
			"error":   "BAD_REQUEST",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "User successfully created",
		"user":    newUser,
	})

}

func GoogleOAuth(c *fiber.Ctx) error {
	fmt.Println(c)
	json := new(contracts.GoogleOAuthRequest)
	if err := c.BodyParser(json); err != nil {
		fmt.Println(err)
		return c.Status(400).JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
			"error":   err.Error(),
		})
	}
	fmt.Println(json.Email)
	fmt.Println("1223")

	token, err := services.HandleGoogleOAuth(*json)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code":    400,
			"message": "Something went wrong",
			"error":   "BAD_REQUEST",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"code":         200,
		"message":      "User successfully logged",
		"access_token": token,
	})

}
