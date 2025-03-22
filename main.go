package main

import (
	"log"
	"net"

	"GoChat/config"
	mygrpc "GoChat/grpc"
	"GoChat/models"
	"GoChat/proto/chat"
	"GoChat/routes"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {

	config.DBInit()
	err := config.DB.AutoMigrate(&models.User{})
	if err != nil {
		panic("migration failed")
	}

	router := gin.Default()

	// add routes
	routes.SetupUserRoutes(router)
	routes.SetupChatRoutes(router)

	// Start gRPC server in a goroutine
	go func() {
		listener, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}

		grpcServer := grpc.NewServer()
		chat.RegisterChatServiceServer(grpcServer, mygrpc.NewChatServer())

		log.Println("gRPC server running on :50051")
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	router.Run("0.0.0.0:8080")
}
