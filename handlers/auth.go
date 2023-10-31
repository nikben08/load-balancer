package handlers

import (
	"fmt"
	"load-balancer/contracts"
	"load-balancer/models"
	encryption "load-balancer/utils/encryption"
	jwt "load-balancer/utils/jwt"

	"github.com/gofiber/fiber/v2"
)

type User models.User

func (h handler) Login(c *fiber.Ctx) error {
	json := new(contracts.LoginRequest)
	if err := c.BodyParser(json); err != nil {
		fmt.Println(err)
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
			"error":   err.Error(),
		})
	}

	hash, err := encryption.GenerateHash(json.Password)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
			"error":   err.Error(),
		})
	}
	user := User{Email: json.Email, Hash: hash}
	found := User{}
	if result := h.DB.Where("email = ?", user.Email).First(&found); result.Error != nil {
		fmt.Println("error")
	}
	if found.Hash == user.Hash {
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

func (h handler) Signup(c *fiber.Ctx) error {
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

	hash, _ := encryption.GenerateHash(json.Password)
	var newUser = User{
		Email:   json.Email,
		Hash:    hash,
		Name:    json.Name,
		Surname: json.Surname,
	}

	if result := h.DB.Create(&newUser); result.Error != nil {
		fmt.Println("error")
	}

	token, _ := jwt.GenerateJwtToken(newUser.Id, json.Email, json.Name)

	var response = contracts.AuthResponse{
		Code:    200,
		Message: "User successfully created",
		Token:   token,
	}

	return c.JSON(response)
}
