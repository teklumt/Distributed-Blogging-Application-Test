package main

import (
	"auth-service/config"
	"auth-service/delivery/routers"

	// "auth_service/infrastructure"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

    // Connect to the database
    config.ConnectDatabase()


    // // Connect to RabbitMQ
    // config.ConnectRabbitMQ()
    // defer config.RabbitMQChannel.Close()

    // Set up the Gin router
    r := gin.Default()

    // Set up routes
    routers.SetupRouter(r)

    // Start the server
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}