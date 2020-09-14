package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8989")
	delErr(err)
	for {
		conn, err := l.Accept()
		delErr(err)
		go handleConn(conn)
	}
}

func delErr(err error) {
	if err != nil {
		log.Fatalf("err:%+v\n", err)
	}
}

func handleConn(conn net.Conn) {
	// header
	var header []byte
	headerLen := 4
	tmp := make([]byte, headerLen-len(header))
	fmt.Print(tmp)
	conn.Read(header)
}
