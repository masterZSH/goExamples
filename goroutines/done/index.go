package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr := "localhost:9999"
	done := make(chan bool)
	go dial(done, addr)
	<-done
}

func dial(done chan bool, addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", conn)
	done <- true
}
