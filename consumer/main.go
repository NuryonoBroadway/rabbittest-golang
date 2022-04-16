package main

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

func main() {
	amqpServerURL := os.Getenv("AMQP_SERVER_URL")

	connectRabbitMQ, err := amqp.Dial(amqpServerURL)

	if err != nil {
		panic(err)
	}
	defer connectRabbitMQ.Close()

	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer channelRabbitMQ.Close()

	messages, err := channelRabbitMQ.Consume(
		"HelloWorld",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println(err)
	}

	log.Println("successfully connected to RabbitMQ")
	log.Println("Waiting for messages")

	forever := make(chan bool)
	log.Println(<-forever)

	go func() {
		for message := range messages {
			log.Printf("> Received message : %s\n", message.Body)
		}
	}()

	<-forever

}
