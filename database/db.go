package database

import (
	"fmt"
	"load-balancer/models"
	"load-balancer/seeds"
	"log"

	"load-balancer/config"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	// Connect to database
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s", config.Config("POSTGRES_HOST"), config.Config("POSTGRES_USER"), config.Config("POSTGRES_PASSWORD"), config.Config("POSTGRES_PORT"))
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err := DB.Exec("DROP DATABASE IF EXISTS users;").Error; err != nil {
		panic(err)
	}

	if err := DB.Exec("CREATE DATABASE users").Error; err != nil {
		panic(err)
	}

	dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", config.Config("POSTGRES_HOST"), config.Config("POSTGRES_USER"), config.Config("POSTGRES_PASSWORD"), config.Config("POSTGRES_DATABASE_NAME"), config.Config("POSTGRES_PORT"))
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Migrate tables
	DB.AutoMigrate(
		&models.User{},
		&models.OAuthProvider{},
		&models.OAuthUser{},
	)

	// Create Super User
	var user = seeds.SuperUser()
	if result := DB.Create(&user); result.Error != nil {
		fmt.Println("Couldn't create super user")
	}

	var oAuthProvider = seeds.OAuthProviders()[0]
	if result := DB.Create(&oAuthProvider); result.Error != nil {
		fmt.Println("Couldn't create super user")
	}

	return DB
}
