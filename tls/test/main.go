package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":10888")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			continue
		}
		handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	content := make([]byte, 1024)
	for {
		n, err := conn.Read(content)
		if n != 0 {
			fmt.Printf("%s", content[:n])
			fmt.Printf("%s", err)
		}
	}
}
