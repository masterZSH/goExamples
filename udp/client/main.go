package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("udp", ":6668")
	if err != nil {
		log.Fatal(err)
	}
	conn.Write([]byte("123"))
}
