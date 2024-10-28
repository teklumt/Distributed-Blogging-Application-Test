package controllers

import (
	"github.com/teklumt/Distributed-Blogging-Application-Test-auth-service/config"
	"github.com/teklumt/Distributed-Blogging-Application-Test-auth-service/domain"
	"github.com/teklumt/Distributed-Blogging-Application-Test-auth-service/messaging"
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
)
type UserController struct {
	UserUsecase domain.UserUsecase
}


func NewAuthController(userUsecase domain.UserUsecase) *UserController {
	return &UserController{UserUsecase: userUsecase}
}


func (uc *UserController) Register(c *gin.Context) {
    var user domain.User

    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    err := uc.UserUsecase.Register(user)
    if err.Message != "" {
        c.JSON(400, gin.H{
            "status": err.StatusCode,
            "message": err.Message,
        })
        return
    }

    rabbitMQConfig, errr := config.NewRabbitMQConfig()
    if errr != nil {
        log.Fatalf("Failed to initialize RabbitMQ: %v", err)
    }
    defer rabbitMQConfig.Close()

    publisher := messaging.NewPublisher(rabbitMQConfig)

    // Create a structured event
    event := messaging.UserRegistrationEvent{
        UserID:   user.ID,
        Message: "Welcome to our platform " + user.Username,
    }
    message, errr := json.Marshal(event)
    if errr != nil {
        log.Printf("Failed to marshal registration event: %v", err)
        return
    }

    errr = publisher.PublishMessage("notification_queue", message)
    if errr != nil {
        log.Printf("Failed to publish registration message: %v", err)
    }

    c.JSON(200, gin.H{
        "status": 200,
        "message": "Account created successfully",
    })
}


func (uc *UserController) Login(c *gin.Context) {
    var user domain.User

    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    result, err := uc.UserUsecase.Login(user.Email, user.Password)
    if err.Message != "" {
        c.JSON(400, gin.H{
            "status": err.StatusCode,
            "message": err.Message,
        })
        return
    }


    rabbitMQConfig, errr := config.NewRabbitMQConfig()
    if errr != nil {
        log.Fatalf("Failed to initialize RabbitMQ: %v", err)
    }
    defer rabbitMQConfig.Close()

    publisher := messaging.NewPublisher(rabbitMQConfig)

    // Create a structured login event
    event := messaging.UserLoginEvent{
        UserID:   result.User.ID,
        Message:"welcome back "  + result.User.Username,
    }
    message, errr := json.Marshal(event)
    if errr != nil {
        log.Printf("Failed to marshal login event: %v", err)
        return
    }

    errr = publisher.PublishMessage("notification_queue", message)
    if errr != nil {
        log.Printf("Failed to publish login message: %v", err)
    }

    c.JSON(200, gin.H{
        "status": 200,
        "data": result,
    })
}
