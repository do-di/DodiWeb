package usecase

import (
	"auth_server/infrastructure/repository"
)

type UserCase struct {
	LoginHandler *LoginHandler
}

func InitUseCase(repository repository.Repository) *UserCase {
	return &UserCase{
		LoginHandler: &LoginHandler{
			Repository: repository,
		},
	}
}
