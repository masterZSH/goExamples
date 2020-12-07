package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
)

func main() {
	cert, err := tls.LoadX509KeyPair("cert/cert.pem", "cert/key.pem")
	if err != nil {
		log.Fatal(err)
	}
	cfg := &tls.Config{Certificates: []tls.Certificate{cert}}
	l, err := tls.Listen("tcp", "localhost:8080", cfg)
	if err != nil {
		log.Fatalf("listen error :%v\n", err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("conn error:%+v\n", err)
			continue
		}
		handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		return
	}
	fmt.Printf("read content %s\n", buf[:n])
}
