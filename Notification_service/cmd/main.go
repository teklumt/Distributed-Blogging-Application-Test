package main

import (
	"encoding/json"
	"log"
	"teklumt/Distributed-Blogging-Application-Test-notification-service/config"
	"teklumt/Distributed-Blogging-Application-Test-notification-service/delivery/routers"
	"teklumt/Distributed-Blogging-Application-Test-notification-service/domain"
	"teklumt/Distributed-Blogging-Application-Test-notification-service/messaging"
	"teklumt/Distributed-Blogging-Application-Test-notification-service/repository"
	"teklumt/Distributed-Blogging-Application-Test-notification-service/usecase"

	"github.com/gin-gonic/gin"
	"github.com/rabbitmq/amqp091-go"
)

func main() {
    // Connect to the database
    config.ConnectDatabase()

    // Initialize RabbitMQ
    rabbitMQConfig, err := config.NewRabbitMQConfig()
    if err != nil {
        log.Fatalf("Failed to initialize RabbitMQ: %v", err)
    }
    defer rabbitMQConfig.Close()

    consumer, errr := messaging.NewConsumer(rabbitMQConfig)
	if errr != nil {
		log.Fatalf("Failed to initialize RabbitMQ consumer: %v", errr)
	}
    notificationRepo := repository.NewNotificationRepository()
    notificationUsecase := usecase.NewNotificationUsecase(notificationRepo)

    // Start RabbitMQ consumer
    go func() {
        err := consumer.ListenToNotifications("notification_queue", func(msg amqp091.Delivery) {
            var notificationEvent messaging.NotificationEvent
            if err := json.Unmarshal(msg.Body, &notificationEvent); err != nil {
                log.Printf("Error parsing message: %v", err)
                return
            }

            notification := &domain.Notification{
                UserID:     notificationEvent.UserID,
                Message:    notificationEvent.Message,
                ReadStatus: false,
            }
            if err := notificationUsecase.CreateNotification(notification); err != nil {
                log.Printf("Error saving notification: %v", err)
            }
        })
        if err != nil {
            log.Fatalf("Error listening to notifications: %v", err)
        }
    }()

    // Set up the Gin router
    r := gin.Default()
    routers.SetupRouter(r)

    // Start the server
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}
