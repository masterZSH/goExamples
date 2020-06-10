package main

import (
	"fmt"
	"os"

	"github.com/Shopify/sarama"
)

func main() {
	kafka := os.Getenv("kafka")
	url := fmt.Sprintf("%s:9095", kafka)
	broker := sarama.NewBroker(url)
	err := broker.Open(nil)
	if err != nil {
		panic(err)
	}

	request := sarama.MetadataRequest{Topics: []string{"myTopic1"}, AllowAutoTopicCreation: true}
	response, err := broker.GetMetadata(&request)
	if err != nil {
		_ = broker.Close()
		panic(err)
	}

	// add topic
	response.AddTopic("myTopic", sarama.ErrNoError)

	response.AddTopicPartition("myTopic", int32(1), int32(1), []int32{1}, []int32{1}, []int32{1}, sarama.ErrNoError)

	fmt.Println("There are", len(response.Topics), "topics active in the cluster.")

	if err = broker.Close(); err != nil {
		panic(err)
	}
}
