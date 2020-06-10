package main
import (
	kafka "github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
)
func main() {
	producer, err := kafka.NewAsyncProducer([]string{"localhost:9095"}, nil)
if err != nil {
    panic(err)
}

defer func() {
    if err := producer.Close(); err != nil {
        log.Fatalln(err)
    }
}()

// Trap SIGINT to trigger a shutdown.
signals := make(chan os.Signal, 1)
signal.Notify(signals, os.Interrupt)

var enqueued, errors int
ProducerLoop:
for {
    select {
    case producer.Input() <- &kafka.ProducerMessage{Topic: "myTopic", Key: nil, Value: kafka.StringEncoder("testing 123")}:
        enqueued++
    case err := <-producer.Errors():
        log.Println("Failed to produce message", err)
        errors++
    case <-signals:
        break ProducerLoop
    }
}

log.Printf("Enqueued: %d; errors: %d\n", enqueued, errors)
}