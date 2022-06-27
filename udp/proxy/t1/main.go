package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("udp", ":3333")
	if err != nil {
		log.Fatal(err)
	}
	conn.Write([]byte("123"))
}
