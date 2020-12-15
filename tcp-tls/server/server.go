package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
)

func main() {
	cert, err := tls.LoadX509KeyPair("./server.pem", "./key.pem")
	if err != nil {
		log.Fatal(err)
	}
	config := &tls.Config{Certificates: []tls.Certificate{cert}}
	l, err := tls.Listen("tcp", ":6366", config)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	buf := make([]byte, 0xfffff)
	for {
		n, err := conn.Read(buf)
		if err != nil || n == 0 {
			return
		}
		fmt.Printf("----------\n")
		fmt.Printf("get %s\n", buf[:n])
	}
}
