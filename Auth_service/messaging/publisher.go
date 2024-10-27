package messaging

import (
	"auth-service/config" // Update this import path if necessary
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	channel *amqp.Channel
}

func NewPublisher(rabbitMQConfig *config.RabbitMQConfig) *Publisher {
	return &Publisher{channel: rabbitMQConfig.Ch}
}

func (p *Publisher) PublishMessage(queueName string, message []byte) error {
	// Declare the queue in case it doesn’t exist
	_, err := p.channel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Printf("Failed to declare a queue: %v", err)
		return err
	}

	err = p.channel.Publish(
		"",        // exchange
		queueName, // routing key (queue name)
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		})
	if err != nil {
		log.Printf("Failed to publish message to queue %s: %v", queueName, err)
		return err
	}

	log.Printf("Message published to queue %s: %s", queueName, message)
	return nil
}


type UserRegistrationEvent struct {
    UserID   uint   `json:"userID"`
    Username string `json:"username"`
}

type UserLoginEvent struct {
    UserID   uint   `json:"userID"`
    Username string `json:"username"`
}


// type UserRegistrationEvent struct {
// 	UserID   uint   `json:"userID"`
// 	Message string `json:"message"`
// }

// type UserLoginEvent struct {
// 	UserID   uint   `json:"userID"`
// 	Username string `json:"username"`
// }

// func (p *Publisher) PublishUserRegistrationEvent(event UserRegistrationEvent) error {
// 	body, err := json.Marshal(event)
// 	if (err != nil) {
// 		return err
// 	}

// 	err = p.channel.Publish(
// 		"user-events",        // exchange
// 		"user.registration",  // routing key
// 		false,                // mandatory
// 		false,                // immediate
// 		amqp.Publishing{
// 			ContentType: "application/json",
// 			Body:        body,
// 		},
// 	)
// 	if err != nil {
// 		log.Printf("Failed to publish registration event: %v", err)
// 		return err
// 	}

// 	log.Println("User registration event published")
// 	return nil
// }

// func (p *Publisher) PublishUserLoginEvent(event UserLoginEvent) error {
// 	body, err := json.Marshal(event)
// 	if err != nil {
// 		return err
// 	}

// 	err = p.channel.Publish(
// 		"user-events", // exchange
// 		"user.login",  // routing key
// 		false,         // mandatory
// 		false,         // immediate
// 		amqp.Publishing{
// 			ContentType: "application/json",
// 			Body:        body,
// 		},
// 	)
// 	if err != nil {
// 		log.Printf("Failed to publish login event: %v", err)
// 		return err
// 	}

// 	log.Println("User login event published")
// 	return nil
// }
