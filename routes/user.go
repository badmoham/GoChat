package routes

import (
	"GoChat/controllers"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine) {
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
}
