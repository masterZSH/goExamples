package main

import (
	"fmt"
	"log"
	"os"

	kafka "github.com/Shopify/sarama"
)

func main() {

	// 2020/06/10 11:42:34 > message sent to partition 0 at offset 2703218
	config := kafka.NewConfig()
	config.Producer.Return.Successes = true
	k := os.Getenv("kafka")
	url := fmt.Sprintf("%s:9095", k)
	producer, err := kafka.NewSyncProducer([]string{url}, nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	msg := &kafka.ProducerMessage{Topic: "myTopic", Value: kafka.StringEncoder("testing 123")}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Printf("FAILED to send message: %s\n", err)
	} else {
		log.Printf("> message sent to partition %d at offset %d\n", partition, offset)
	}

}
