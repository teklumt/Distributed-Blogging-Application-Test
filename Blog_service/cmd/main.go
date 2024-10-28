package main

import (
	"github.com/teklumt/Distributed-Blogging-Application-Test-blog-service/config"
	"github.com/teklumt/Distributed-Blogging-Application-Test-blog-service/delivery/routers"
	"github.com/teklumt/Distributed-Blogging-Application-Test-blog-service/messaging"

	// "auth_service/infrastructure"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

    // Connect to the database
    config.ConnectDatabase()


     // Set up the Gin router    rabbitMQConfig, err := config.NewRabbitMQConfig()
    rabbitMQConfig, err := config.NewRabbitMQConfig()
     if err != nil {
         log.Fatalf("Failed to initialize RabbitMQ: %v", err)
     }
    defer rabbitMQConfig.Close()
 
    publisher :=  messaging.NewPublisher(rabbitMQConfig)

    
    
    r := gin.Default()

    // Set up routes
    routers.SetupRouter(r, publisher)

    // Start the server
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}