package controllers

import (
	"net/http"

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
