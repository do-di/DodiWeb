package routes

import (
	"auth_server/web/controller"

	"github.com/gin-gonic/gin"
)

func InitAuthRoutes(router *gin.Engine, controller *controller.Controller) {
	authRouter := router.Group("auth")
	authRouter.POST("/login", controller.AuthController.Login)
}
