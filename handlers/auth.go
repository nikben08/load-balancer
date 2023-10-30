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
type UserClaim models.UserClaim
type PhoneNumber models.PhoneNumber
type Address models.Address

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

	user := User{Email: json.Email, Hash: encryption.GenerateHash(json.Password)}
	found := User{}
	if result := h.DB.Where("email = ?", user.Email).First(&found); result.Error != nil {
		fmt.Println("error")
	}
	if found.Hash == user.Hash {
		var userClaim models.UserClaim
		if result := h.DB.Where("user_id = ?", user.ID).First(&userClaim); result.Error != nil {
			fmt.Println("error")
		}

		token, _ := jwt.GenerateJwtToken(found.ID, found.Email, found.Name, userClaim.AccessLevel)

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

	var newUser = User{
		Email:   json.Email,
		Hash:    encryption.GenerateHash(json.Password),
		Name:    json.Name,
		Sername: json.Sername,
	}

	if result := h.DB.Create(&newUser); result.Error != nil {
		fmt.Println("error")
	}

	token, _ := jwt.GenerateJwtToken(newUser.ID, json.Email, json.Name, 0)

	var response = contracts.AuthResponse{
		Code:    200,
		Message: "User successfully created",
		Token:   token,
	}

	return c.JSON(response)
}
