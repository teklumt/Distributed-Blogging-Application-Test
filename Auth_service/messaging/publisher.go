package messaging

import (
	"auth-service/config"
	"encoding/json"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type UserRegistrationEvent struct {
    UserID   uint `json:"userID"`
    Username string `json:"username"`
}

type UserLoginEvent struct {
	UserID   uint `json:"userID"`
	Username string `json:"username"`
}


func PublishUserRegistrationEvent(event UserRegistrationEvent) error {
    body, err := json.Marshal(event)
    if err != nil {
        return err
    }

    err = config.RabbitMQChannel.Publish(
        "user-events",        // exchange
        "user.registration",   // routing key
        false,                 // mandatory
        false,                 // immediate
        amqp091.Publishing{
            ContentType: "application/json",
            Body:        body,
        },
    )
    if err != nil {
        log.Printf("Failed to publish registration event: %v", err)
        return err
    }

    log.Println("User registration event published")
    return nil
}


func PublishUserLoginEvent(event UserLoginEvent) error {
	body, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = config.RabbitMQChannel.Publish(
		"user-events",        // exchange
		"user.login",          // routing key
		false,                 // mandatory
		false,                 // immediate
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Printf("Failed to publish login event: %v", err)
		return err
	}

	log.Println("User login event published")
	return nil
}	