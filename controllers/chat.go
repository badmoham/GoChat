package controllers

import (
	"net/http"
	"strconv"

	"GoChat/services"
	"github.com/gin-gonic/gin"
)

func StartPersonalChatRoom(cnx *gin.Context) {
	var userInput struct {
		JWTToken                       string `json:"jwt_token" binding:"required"`
		DestinationChatUserPhoneNumber string `json:"destination_chat_user_phone_number" binding:"required"`
	}

	if err := cnx.BindJSON(userInput); err != nil {
		return
	}
	resolvedSourceUserID, err := services.DecryptJWT(userInput.JWTToken)
	if err != nil {
		cnx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	resolvedDestUserID, err := services.ResolveUserIDFromPhoneNumber(userInput.DestinationChatUserPhoneNumber)
	if err != nil {
		cnx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	newChatID, err := services.CreateNewPersonalChat(resolvedSourceUserID, resolvedDestUserID)
	if err != nil {
		cnx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	cnx.JSON(http.StatusCreated, gin.H{"message": "conversation started", "room_id": newChatID})
}

func GetAllMyChat(cnx *gin.Context) {
	var userInput struct {
		JWTToken string `json:"jwt_token" binding:"required"`
	}
	if err := cnx.BindJSON(userInput); err != nil {
		return
	}
	resolvedSourceUserID, err := services.DecryptJWT(userInput.JWTToken)
	if err != nil {
		cnx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	allChats, err := services.GetUserAllChats(resolvedSourceUserID)
	if err != nil {
		cnx.JSON(http.StatusServiceUnavailable, gin.H{"message": "query ran into an error"})
	}
	cnx.JSON(http.StatusOK, allChats)

}

func SendTextMessage(cnx *gin.Context) {
	var userInput struct {
		JWTToken   string `json:"jwt_token" binding:"required"`
		Message    string `json:"message" binding:"required"`
		ChatRoomID uint   `json:"chat_room_id" binding:" required"`
	}
	if err := cnx.BindJSON(userInput); err != nil {
		return
	}
	resolvedSourceUserID, err := services.DecryptJWT(userInput.JWTToken)
	if err != nil {
		cnx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	err = services.SubmitTextMessage(resolvedSourceUserID, userInput.ChatRoomID, userInput.Message)
	if err != nil {
		cnx.JSON(http.StatusUnauthorized, err.Error())
		return
	}
	cnx.JSON(http.StatusAccepted, gin.H{"message": "message was sent successfully"})

}

func GetChatRoomMessages(cnx *gin.Context) {
	// Parse query parameters
	JWTToken := cnx.DefaultQuery("jwt_token", "")
	chatID, err := strconv.ParseUint(cnx.Param("chat_id"), 10, 64)
	if err != nil {
		cnx.JSON(http.StatusBadRequest, gin.H{"message": "chat_id was not provided correctly"})
		return
	}
	page, err := strconv.Atoi(cnx.DefaultQuery("page", "1"))
	if err != nil {
		cnx.JSON(http.StatusBadRequest, gin.H{"message": "page was not provided correctly"})
		return
	}
	limit, err := strconv.Atoi(cnx.DefaultQuery("limit", "10"))
	if err != nil || limit > 100 {
		cnx.JSON(http.StatusBadRequest, gin.H{"message": "limit was not provided correctly"})
		return
	}
	offset := (page - 1) * limit
	resolvedSourceUserID, err := services.DecryptJWT(JWTToken)
	if err != nil {
		cnx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	result, err := services.GetUserChatRoomMessages(resolvedSourceUserID, uint(chatID), limit, offset)
	if err != nil {
		cnx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	cnx.JSON(http.StatusOK, result)

}
