package config

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQConfig struct {
    Conn *amqp.Connection
    Ch   *amqp.Channel
}

func NewRabbitMQConfig() (*RabbitMQConfig, error) {
    // conn, err := amqp.Dial("amqp://admin:password@localhost:5672/")
    // conn, err := amqp.Dial("amqp://admin:password@rabbitmq:5672/")
    conn, err := amqp.Dial("amqp://admin:password@rabbitmq:5672/")

    if err != nil {
        log.Fatalf("Failed to connect to RabbitMQ: %v", err)
        return nil, err
    }

    ch, err := conn.Channel()
    if err != nil {
        log.Fatalf("Failed to open a channel: %v", err)
        return nil, err
    }

    return &RabbitMQConfig{
        Conn: conn,
        Ch:   ch,
    }, nil
}

func (r *RabbitMQConfig) Close() {
    r.Ch.Close()
    r.Conn.Close()
}
