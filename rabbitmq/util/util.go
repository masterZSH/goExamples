package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

var (
	in  = os.Stdin
	out = os.Stdout

	// host  => rabbitmq host
	// queue => rabbitmq queue name
	// user  => rabbitmq user
	// pwd   => rabbitmq password
	host, queue, user, pwd string
	// port  => rabbitmq port
	port int
)

func init() {
	flag.StringVar(&host, "h", "", "rabbitmq server host")
	flag.StringVar(&queue, "q", "", "rabbitmq queue name")
	flag.IntVar(&port, "p", 5672, "rabbitmq port")
	flag.StringVar(&user, "u", "", "rabbitmq username")
	flag.StringVar(&pwd, "d", "", "rabbitmq password")
}

func main() {
	flag.Parse()
	if host == "" {
		log.Fatal("miss param host")
	}
	if queue == "" {
		log.Fatal("miss param queue name")
	}
	if user == "" {
		log.Fatal("miss param username")
	}
	if pwd == "" {
		log.Fatal("miss param password")
	}
	conn := newConn()
	channel := newChannel(conn)
	done := make(chan bool)
	msgs := make(chan string, 10)
	go scanInput(done, msgs)
	go publish(msgs, channel)
	<-done
}

// 扫描输入
func scanInput(ch chan bool, msgs chan string) {
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		msgs <- scanner.Text()
	}
	ch <- true
}

func logError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func buildURL() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d/", user, pwd, host, port)
}

func newConn() *amqp.Connection {
	url := buildURL()
	conn, err := amqp.Dial(url)
	logError(err)
	fmt.Printf("conn to %s success\n", host)
	return conn
}

func newChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	logError(err)
	return ch
}

func newQueue(ch *amqp.Channel) amqp.Queue {
	q, err := ch.QueueDeclare(
		queue, // queue name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	logError(err)
	return q
}

func newQueueByConn(conn *amqp.Connection) amqp.Queue {
	ch := newChannel(conn)
	q, err := ch.QueueDeclare(
		queue, // queue name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	logError(err)
	return q
}

func publish(msgs chan string, channel *amqp.Channel) {
	queue := newQueue(channel)
	for msg := range msgs {
		err := channel.Publish(
			"",         // exchange
			queue.Name, // routing key
			false,      // mandatory
			false,      // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(msg),
			})
		if err != nil {
			fmt.Printf("send msg error :%v\n", err)
		}
	}
}
