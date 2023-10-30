package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model
	Id uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();"`
}
