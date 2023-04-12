package usecase

import (
	"fmt"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/utils/security"

	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/repository"
)

type AuthenticationUseCase interface {
	Login(username string, password string) (string, error)
}

type authenticationUseCase struct {
	repo         repository.UserRepository
	tokenService security.AccessToken
}

func (a *authenticationUseCase) Login(username string, password string) (string, error) {
	user, err := a.repo.GetByUsernamePassword(username, password)
	var token string
	if err != nil {
		return "", fmt.Errorf("user with username: %s not found", username)
	}
	if user != nil {
		token, err = a.tokenService.CreateAccessToken(user)
		fmt.Println("token:", token)
		if err != nil {
			return "", err
		}
	}
	return token, nil
}

func NewAuthenticationUseCase(repo repository.UserRepository, tokenService security.AccessToken) AuthenticationUseCase {
	return &authenticationUseCase{repo: repo, tokenService: tokenService}
}
