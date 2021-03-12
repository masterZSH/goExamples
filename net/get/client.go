package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}

	go readConn(conn)
	go writeConn(conn)

	select {}
}

func readConn(conn net.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			return
		}

		fmt.Printf("%s\n", buf[:n])
	}
}

func writeConn(conn net.Conn) {
	conn.Write([]byte("gettest"))
}
