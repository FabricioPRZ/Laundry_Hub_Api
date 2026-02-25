package application

import (
	"laundry-hub-api/src/user/domain"
	"laundry-hub-api/src/user/domain/entities"
)

type OAuthService struct {
	userRepo domain.IUserRepository
}

func NewOAuthService(userRepo domain.IUserRepository) *OAuthService {
	return &OAuthService{userRepo: userRepo}
}

func (os *OAuthService) FindOrCreateOAuthUser(email, provider, oauthID, name string) (*entities.User, error) {
	user, err := os.userRepo.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	if user != nil {
		return user, nil
	}

	newUser := &entities.User{
		Name:          name,
		Email:         email,
		OAuthProvider: provider,
		OAuthID:       &oauthID,
		Role:          "USER",
	}

	return os.userRepo.Save(newUser)
}
