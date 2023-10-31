package routers

import (
	"proxolab-ecommerce-authorization/handlers"

	"github.com/gofiber/fiber/v2"
)

func Initalize(app *fiber.App) {
	api := app.Group("/api/v1")
	auth := api.Group("/auth")
	auth.Post("/signup", handlers.Signup)
	auth.Post("/login", handlers.Login)

	oauth := api.Group("/oauth")
	oauth.Get("/google", handlers.GoogleOAuth)
}
