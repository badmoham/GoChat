package controllers

import (
	"net/http"

	"GoChat/models"
	"GoChat/services"
	"github.com/gin-gonic/gin"
)

func Register(cnx *gin.Context) {
	var userInput struct {
		Name        string `json:"name" binding:"required"`
		PhoneNumber string `json:"phone_number" binding:"required"`
		Password    string `json:"password" binding:"required"`
	}

	if err := cnx.BindJSON(userInput); err != nil {
		return
	}
	newUser := &models.User{
		Name:        userInput.Name,
		PhoneNumber: userInput.PhoneNumber,
		Password:    userInput.Password,
	}
	if err := services.CreateUser(newUser); err != nil {
		cnx.JSON(http.StatusInternalServerError, gin.H{"code": 6001})
		return
	}
	cnx.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}

func Login(cnx *gin.Context) {
	var userInput struct {
		PhoneNumber string `json:"phone_number" binding:"required"`
		Password    string `json:"password" binding:"required"`
	}
	if err := cnx.BindJSON(userInput); err != nil {
		return
	}
	JWT, err := services.AuthenticateUser(userInput.PhoneNumber, userInput.Password)
	if err != nil {
		cnx.JSON(http.StatusNotFound, gin.H{"message": err})
		return
	}
	cnx.JSON(http.StatusOK, gin.H{"token": JWT})
}
