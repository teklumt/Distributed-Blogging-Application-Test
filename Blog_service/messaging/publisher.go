package messaging

import (
	"auth-service/config"
	"encoding/json"
	"log"

	"github.com/rabbitmq/amqp091-go"
)





type NewBlogPostedEvent struct {
	UserID string `json:"user_id"`
	Message string `json:"message"`
}


func PublishNewBlogPostedEvent(event NewBlogPostedEvent) error {
	body, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = config.RabbitMQChannel.Publish(
		"blog-events",        // exchange
		"blog.posted", // routing key
		false,                 // mandatory
		false,                 // immediate
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Printf("Failed to publish new blog posted event: %v", err)
		return err
	}

	log.Println("New blog posted event published")
	return nil
	

}


type BlogDeletedEvent struct {
	UserId string `json:"user_id"`
	Message string `json:"message"`

}

func PublishBlogDeletedEvent(event BlogDeletedEvent) error {
	body, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = config.RabbitMQChannel.Publish(
		"blog-events",        // exchange
		"blog.deleted", // routing key
		false,                 // mandatory
		false,                 // immediate
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Printf("Failed to publish blog deleted event: %v", err)
		return err
	}

	log.Println("Blog deleted event published")
	return nil
}




type CommentPostedEvent struct {
	Message string `json:"message"`
	PostID  uint `json:"post_Id"`
}


func PublishCommentPostedEvent(event CommentPostedEvent) error {
	body, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = config.RabbitMQChannel.Publish(
		"comment-events",        // exchange
		"comment.posted", // routing key
		false,                 // mandatory
		false,                 // immediate
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Printf("Failed to publish comment posted event: %v", err)
		return err
	}

	log.Println("Comment posted event published")
	return nil
}