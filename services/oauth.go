package services

import (
	"load-balancer/contracts"
	"load-balancer/models"
	"load-balancer/repositories"
	jwt "load-balancer/utils/jwt"

	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

var (
	googleOauthConfig *oauth2.Config
	// TODO: randomize it
	oauthStateString = "pseudo-random"
)

func HandleGoogleOAuth(oAuthUser contracts.GoogleOAuthRequest) (string, error) {
	isOAuthUserExists, _ := repositories.WhetherUserExistInOAuth(oAuthUser.Id)
	isUserExists, user := repositories.WhetherUserExistInUsers(oAuthUser.Email)

	if !isOAuthUserExists && !isUserExists {
		user := &models.User{
			Email:   oAuthUser.Email,
			Name:    oAuthUser.Name,
			Surname: oAuthUser.Surname,
		}

		err := repositories.CreateNewUser(user)
		if err != nil {
			return "", err
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

	err := repositories.GetUserByEmail(&user)
	if err != nil {
		return "", err
	}

	token, err := jwt.GenerateJwtToken(user.Id, user.Email, user.Name)
	if err != nil {
		return "", err
	}
	return token, nil
}
