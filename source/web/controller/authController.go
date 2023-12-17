package controller

import (
	"auth_server/application/usecase"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	UserCase *usecase.UserCase
}

func (a AuthController) Login(c *gin.Context) {
	command := usecase.LoginCommand{
		Username: "abc@gmail.com",
		Password: "123@bcd",
	}
	response, _ := a.UserCase.LoginHandler.Handle(command)
	c.JSON(200,
		response,
	)
}
