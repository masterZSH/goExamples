package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/streadway/amqp"
)

var (
	host = os.Getenv("rbhost")
	user = os.Getenv("rbuser")
	pwd  = os.Getenv("rbpwd")
)

func main() {
	url := fmt.Sprintf("amqp://%s:%s@%s:5672/", user, pwd, host)
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// exchange
	err = ch.ExchangeDeclare(
		"test_sms", // name
		"fanout",   // type
		false,      // durable
		false,      // auto-deleted
		false,      // internal
		false,      // no-wait
		nil,        // arguments
	)

	body := "测试短信"
	err = ch.Publish(
		"test_sms", // exchange
		"",         // routing key
		false,      // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})

	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
	time.Sleep(time.Second)

}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
