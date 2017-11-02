package queue

import (
	"github.com/geertvl/amqptest/error"
	"github.com/streadway/amqp"
	"log"
)

func Connect() *amqp.Connection {
	conn, err := amqp.Dial("amqp://rabbitmq:rabbitmq@localhost:5672")
	error.FailOnError(err, "Failed to connect to RabbitMQ")

	return conn
}

func Declare(ch *amqp.Channel, exchange string) {
	err := ch.ExchangeDeclare(
		exchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	error.FailOnError(err, "Failed to declare an exchange")
}

func Publish(ch *amqp.Channel, exchange string, message []byte) {
	err := ch.Publish(
		exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		})
	error.FailOnError(err, "Failed to publish a message")
	log.Printf("Sent: %s", message)
}
