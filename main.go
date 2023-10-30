package main
<<<<<<< HEAD

import (
	"log"
	"os"
	"load-balancer/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))
	routers.Initalize(app)
	log.Fatal(app.Listen(":" + getenv("PORT", "83")))
}
=======
>>>>>>> 3e42eaab96636658b0510040657ebdb968aa0125
