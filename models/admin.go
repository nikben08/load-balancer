package models

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	ID          uint `json:"id" gorm:"primaryKey"`
	UserID      uint `json:"user_id" gorm:"primaryKey"`
	AccessLevel int  `json:"access_level" gorm:"access_level"`
}
