package receive

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		fmt.Printf("error : %s  message : %s ", err, msg)
		log.Fatalf("%s:%s\n", err, msg)
	}

}

func Consumer() {
	fmt.Println("Consumer function body")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to rabbitmq")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"alert", //name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no_wait
		nil,     // arguments
	)

	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	failOnError(err, "failed out to registered consumer ")

	var forever = make(chan struct{})

	go func() {
		for d := range msgs {
			log.Printf("Received a message : %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for message . To exit press CTRL+C")

	<-forever
}
