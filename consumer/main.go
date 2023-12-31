package main

import (
	"context"
	"fmt"
	"github.com/rabbitmq/amqp091-go"
)

func main() {
	connection, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	// Create channel
	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}

	// Consume queue
	ctx := context.Background()
	emailConsumer, err := channel.ConsumeWithContext(ctx, "email", "email-consumer", true, false, false, false, nil)
	if err != nil {
		panic(err)

	}

	// read data consumer
	for message := range emailConsumer {
		fmt.Println("Routing key: ", message.RoutingKey)
		fmt.Println("Body: ", string(message.Body))
	}
}
