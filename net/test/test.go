package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		log.Fatalf("net.Listen(tcp, localhost:0): %v", err)
	}
	fmt.Println(l.Addr().String())
}
