package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetPostgresDSN() string {
	return "host=localhost user=postgres password=0811 dbname=users port=5432 sslmode=disable"
}

func GetJwtSecretKey() string {
	return os.Getenv("JWT_SECRET_KEY")
}

func GetJwtExpirationMinutes() string {
	return os.Getenv("JWT_EXPIRATION_MINUTES")
}

// Config func to get env value from key ---
func Config(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}
