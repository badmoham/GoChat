package main

import (
	"GoChat/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.SetupUserRoutes(router)
	router.Run("0.0.0.0:8080")
}
