package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", sayHello)

	router.Run("0.0.0.0:8080")
}

func sayHello(context *gin.Context) {
	context.JSON(http.StatusOK, "hello world!")
}
