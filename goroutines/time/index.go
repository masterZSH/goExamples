package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:9999")
	handleError(err)
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, time.Now().Format("15:04:05.00000\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func handleError(err error) {
	if err != nil {
		fmt.Printf("error:%+v\n", err)
		os.Exit(1)
	}
}
