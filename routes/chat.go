package routes

import (
	"GoChat/controllers"
	"github.com/gin-gonic/gin"
)

func SetupChatRoutes(router *gin.Engine) {
	router.POST("/create_personal_chat_room", controllers.StartPersonalChatRoom)
}
