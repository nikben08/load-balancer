package repositories

import (
	"load-balancer/models"

	"github.com/google/uuid"
)

func CreateUser(user models.User) (models.User, error) {
	if result := DB.Create(&user); result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func CreateOAuthUser(oauth models.OAuthUser) (models.OAuthUser, error) {
	if result := DB.Create(&oauth); result.Error != nil {
		return models.OAuthUser{}, result.Error
	}
	return oauth, nil
}

func GetUserByEmail(user *models.User) error {
	if result := DB.Where("email = ?", user.Email).First(&user); result.Error != nil {
		return result.Error
	}
	return nil
}

func WhetherUserExistInOAuth(oauthUserID string) (bool, models.OAuthUser) {
	var oauthUser = models.OAuthUser{}
	if result := DB.Where("oauth_user_id = ?", oauthUserID).First(&oauthUser); result.Error != nil {
		return false, oauthUser
	}
	return true, oauthUser
}

func WhetherUserExistInUsers(email string) (bool, models.User) {
	var user = models.User{}
	if result := DB.Where("email = ?", email).First(&user); result.Error != nil {
		return false, user
	}
	return true, user
}

func GetOAuthProviderByName(providerName string) (uuid.UUID, error) {
	var provider = models.OAuthProvider{}
	if result := DB.Where("name = ?", providerName).First(&provider); result.Error != nil {
		return provider.Id, result.Error
	}
	return provider.Id, nil
}
