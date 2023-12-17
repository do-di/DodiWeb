package usecase

import (
	repositoryinterface "auth_server/application/interface/repositoryInterface"
	"fmt"
)

type LoginHandler struct {
	Repository repositoryinterface.RepositoryInterface
}

type LoginCommand struct {
	Username string
	Password string
}

type LoginResponse struct {
	Token string
}

func (s *LoginHandler) Handle(request LoginCommand) (*LoginResponse, error) {
	test := s.Repository.UserRepositoryInterface.GetUserByUserName("abcd")
	fmt.Println(test)
	return &LoginResponse{
		Token: "abcd",
	}, nil
}
