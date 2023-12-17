package routes

import (
	"auth_server/web/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine, controller *controller.Controller) {
	InitAuthRoutes(router, controller)
}
