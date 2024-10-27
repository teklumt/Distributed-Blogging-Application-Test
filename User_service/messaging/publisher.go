package messaging

import (
	"auth-service/config"
	"encoding/json"
	"log"

	"github.com/rabbitmq/amqp091-go"
)


type UserAccountUpdateEvent struct {
	UserID   uint `json:"userID"`
	Username string `json:"username"`
}


func PublishUserAccountUpdateEvent(event UserAccountUpdateEvent) error {
	body, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = config.RabbitMQChannel.Publish(
		"user-events",        // exchange
		"user.account.update", // routing key
		false,                 // mandatory
		false,                 // immediate
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Printf("Failed to publish account update event: %v", err)
		return err
	}

	log.Println("User account update event published")
	return nil
}