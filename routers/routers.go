package routers

import (
	"proxolab-ecommerce-authorization/handlers"

	"github.com/gofiber/fiber/v2"
)

<<<<<<< HEAD
func Initalize(app *fiber.App) {
	api := app.Group("/api/v1")
	auth := api.Group("/auth")
	auth.Post("/signup", handlers.Signup)
	auth.Post("/login", handlers.Login)

	oauth := api.Group("/oauth")
	oauth.Get("/google", handlers.GoogleOAuth)
=======
func SetupRoutes(app *fiber.App, DB *gorm.DB) {
	h := handlers.New(DB)
	api := app.Group("/api/v1")
	auth := api.Group("/auth")
	auth.Post("/signup", h.Signup)
	auth.Post("/login", h.Login)

>>>>>>> 3e42eaab96636658b0510040657ebdb968aa0125
}
