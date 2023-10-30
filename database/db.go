package database

import (
	"fmt"
	"log"
	"proxolab-ecommerce-authorization/models"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"honnef.co/go/tools/config"
)

func Init() *gorm.DB {
	// Connect to database
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s", config.Config("DBHost"), config.Config("DBUsername"), config.Config("DBUserPassword"), config.Config("DBPort"))
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

	dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", config.Config("DBHost"), config.Config("DBUsername"), config.Config("DBUserPassword"), config.Config("DBName"), config.Config("DBPort"))
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Migrate tables
	DB.AutoMigrate(&models.User{}, &models.Admin{})

	// Create Super User
	var user = &models.User{
		Email: "admin@admin.com",
		Hash:  "08f46d939d7ff7ebb1df9de1cc246135c7a8694cabf0e1c1cc67c50e12832f08",
	}

	if result := DB.Create(&user); result.Error != nil {
		fmt.Println("Couldn't create super user")
	}

	var admin = &models.Admin{
		UserID:      user.ID,
		AccessLevel: 1,
	}

	if result := DB.Create(&admin); result.Error != nil {
		fmt.Println("Couldn't create super user")
	}

	return DB
}
