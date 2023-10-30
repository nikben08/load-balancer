package models

import (
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	UserID  uint   `json:"user_id" gorm:"user_id"`
	ID      int    `json:"id" gorm:"primaryKey"`
	Address string `json:"address" gorm:"address"`
}
