package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uint          `json:"id" gorm:"primaryKey"`
	Hash        string        `json:"hash" gorm:"hash"`
	Email       string        `json:"email" gorm:"email"`
	Name        string        `json:"name" gorm:"name"`
	Sername     string        `json:"sername" gorm:"sername"`
	PhoneNumber []PhoneNumber `gorm:"foreignKey:UserID; constraint:OnDelete:CASCADE;" json:"-"`
	Address     []Address     `gorm:"foreignKey:UserID; constraint:OnDelete:CASCADE;" json:"-"`
}
