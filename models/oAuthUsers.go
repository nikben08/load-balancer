package models

import (
	"github.com/google/uuid"
)

type OAuthUser struct {
	BaseModel
	UserId      uuid.UUID `gorm:"user_id"`
	OAuthUserId uuid.UUID `gorm:"oauth_user_id"`
	ProviderId  uuid.UUID `gorm:"provider_id"`
}
