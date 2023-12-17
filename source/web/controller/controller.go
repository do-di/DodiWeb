package controller

import "auth_server/application/usecase"

type Controller struct {
	AuthController *AuthController
}

func InitController(usecase *usecase.UserCase) *Controller {
	return &Controller{
		AuthController: &AuthController{
			UserCase: usecase,
		},
	}
}
