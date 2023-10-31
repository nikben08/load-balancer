package seeds

import (
	"load-balancer/models"
	encryption "load-balancer/utils/encryption"
)

func OAuthProviders() []models.OAuthProvider {
	var oAuthProviders = []models.OAuthProvider{
		models.OAuthProvider{
			Name: "Google",
		},
	}
	return oAuthProviders
}

func SuperUser() models.User {
	password := "08112001"
	hash, _ := encryption.GenerateHash([]byte(password))
	var superUser = models.User{
		Email:   "admin@gmail.com",
		Hash:    hash,
		Name:    "Tony",
		Surname: "Ferguson",
	}
	return superUser
}
