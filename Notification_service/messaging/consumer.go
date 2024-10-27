package messaging

import (
	"fmt"

	"github.com/teklumt/Distributed-Blogging-Application-Test-notification-service/config"

	amqp "github.com/rabbitmq/amqp091-go"
)

type NotificationEvent struct {
	UserID  uint   `json:"user_id"`
	Message string `json:"message"`
}

type Consumer struct {
	channel *amqp.Channel
}

func NewConsumer(cfg *config.RabbitMQConfig) (*Consumer, error) {
	return &Consumer{channel: cfg.Ch}, nil
}

func (c *Consumer) ListenToNotifications(queueName string, handleFunc func(msg amqp.Delivery)) error {
    // Ensure the queue is declared
    _, err := c.channel.QueueDeclare(
        queueName,
        true,  // durable
        false, // auto-delete
        false, // exclusive
        false, // no-wait
        nil,   // args
    )
    if err != nil {
        return fmt.Errorf("failed to declare queue %s: %v", queueName, err)
    }

    msgs, err := c.channel.Consume(
        queueName,
        "",
        true,
        false,
        false,
        false,
        nil,
    )
    if err != nil {
        return err
    }

    for msg := range msgs {
        handleFunc(msg)
    }

    return nil
}
