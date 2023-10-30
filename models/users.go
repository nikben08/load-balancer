package models

import (
	"github.com/google/uuid"
)

type User struct {
	BaseModel
	Hash    string      `gorm:"hash"`
	Email   string      `gorm:"email"`
	Name    string      `gorm:"name"`
	Surname string      `gorm:"surname"`
	OAuth   []OAuthUser `gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE;" json:"-"`
	Claim   uuid.UUID   `gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE;" json:"-"`
}
