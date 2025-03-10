package routes

import (
	"GoChat/controllers"
	"github.com/gin-gonic/gin"
)

func SetupChatRoutes(router *gin.Engine) {
	router.POST("/create_personal_chat_room", controllers.StartPersonalChatRoom)
	router.POST("/get_all_chats", controllers.GetAllMyChat)
	router.POST("/send_text_message", controllers.SendTextMessage)
	router.POST("/get_chat_message/:chat_ids", controllers.GetChatRoomMessages)
}
