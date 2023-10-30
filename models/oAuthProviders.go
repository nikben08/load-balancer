package models

type OAuthProvider struct {
	BaseModel
	Name string `gorm:"name"`
}
