package main

import (
	"log"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalln("failed to open connection", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalln("failed to open a channel", err)
	}
	defer ch.Close()

	// declaring a queue is idempotent, can be done
	q, err := ch.QueueDeclare(
		"my-queue-name", // name
		false,           // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	if err != nil {
		log.Fatalln("failed to declare a queue", err)
	}

	messageBody := "This_is_a_message " + time.Now().String()
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(messageBody),
		},
	)
	if err != nil {
		log.Fatalln("failed to publish a message", err)
	}
}
