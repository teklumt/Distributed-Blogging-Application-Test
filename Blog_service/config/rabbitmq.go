package config

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

var RabbitMQChannel *amqp091.Channel

func ConnectRabbitMQ() {
    conn, err := amqp091.Dial("amqp://Apirabitmq:Apirabitmq@rabbitmq.rabbitmq.svc.cluster.local:5672/")
    if err != nil {
        log.Fatalf("Failed to connect to RabbitMQ: %v", err)
    }

    ch, err := conn.Channel()
    if err != nil {
        log.Fatalf("Failed to open a channel: %v", err)
    }

    RabbitMQChannel = ch
    log.Println("Connected to RabbitMQ")
}
