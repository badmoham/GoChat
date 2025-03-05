package main

import (
	"GoChat/config"
	"GoChat/models"
	"GoChat/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	config.DBInit()
	err := config.DB.AutoMigrate(&models.User{})
	if err != nil {
		panic("migration failed")
	}

	router := gin.Default()

	routes.SetupUserRoutes(router)
	router.Run("0.0.0.0:8080")
}
