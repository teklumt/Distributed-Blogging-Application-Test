package config

import (
	"log"

	 "github.com/rabbitmq/amqp091-go"
	"github.com/streadway/amqp"
)

var RabbitMQChannel *amqp091.Channel

func ConnectRabbitMQ() {
    // conn, err := amqp.Dial("amqp://admin:password@127.0.0.1:5672/")

    // if err != nil {
    //     log.Fatalf("Failed to connect to RabbitMQ: %v", err)
    // }

    // ch, err := conn.Channel()
    // if err != nil {
    //     log.Fatalf("Failed to open a channel: %v", err)
    // }

    // RabbitMQChannel = ch
    // log.Println("Connected to RabbitMQ")
}
