package models

import (
	"gorm.io/gorm"
)

type PhoneNumber struct {
	gorm.Model
	UserID      uint   `json:"user_id" gorm:"user_id"`
	ID          int    `json:"id" gorm:"primaryKey"`
	PhoneNumber string `json:"phone_number" gorm:"phone_number"`
}
