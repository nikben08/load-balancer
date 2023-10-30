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
	var superUser = models.User{
		Email:   "admin@gmail.com",
		Hash:    encryption.GenerateHash("08112001"),
		Name:    "Tony",
		Surname: "Ferguson",
	}
	return superUser
}
