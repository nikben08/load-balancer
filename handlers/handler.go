package handlers

import (
	"proxolab-ecommerce-authorization/models"

	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) handler {
	return handler{db}
}

type User models.User
