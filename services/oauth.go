package services

import (
	"load-balancer/contracts"
	"load-balancer/models"
	"load-balancer/repositories"

	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

var (
	googleOauthConfig *oauth2.Config
	// TODO: randomize it
	oauthStateString = "pseudo-random"
)

func HandleGoogleOAuth(oAuthUser contracts.GoogleOAuthRequest) string {
	isOAuthUserExists, _ := repositories.WhetherUserExistInOAuth(oAuthUser.Id)
	isUserExists, user := repositories.WhetherUserExistInUsers(oAuthUser.Email)
	token := "3333"

	if !isOAuthUserExists && !isUserExists {
		user := models.User{
			Email:   oAuthUser.Email,
			Name:    oAuthUser.Name,
			Surname: oAuthUser.Surname,
		}

		user, err := repositories.CreateUser(user)
		if err != nil {
			return err.Error()
		}

		oAuthProviderId, _ := repositories.GetOAuthProviderByName("Google")
		uuid, _ := uuid.Parse(oAuthUser.Id)
		repositories.CreateOAuthUser(models.OAuthUser{
			UserId:      user.Id,
			OAuthUserId: uuid,
			ProviderId:  oAuthProviderId,
		})
	}

	if !isOAuthUserExists && isUserExists {
		oAuthProviderId, _ := repositories.GetOAuthProviderByName("Google")
		uuid, _ := uuid.Parse(oAuthUser.Id)
		repositories.CreateOAuthUser(models.OAuthUser{
			UserId:      user.Id,
			OAuthUserId: uuid,
			ProviderId:  oAuthProviderId,
		})
	}

	return token
}
