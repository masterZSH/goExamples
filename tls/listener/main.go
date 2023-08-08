package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatal(err)
	}
	serveTls(l)
}

func serveTls(l net.Listener) {
	cert, err := tls.LoadX509KeyPair("cert/server.pem", "cert/key.pem")
	if err != nil {
		log.Fatal(err)
	}
	cfg := &tls.Config{Certificates: []tls.Certificate{cert}}
	nl := tls.NewListener(l, cfg)
	serve(nl)
}

func serve(l net.Listener) {
	for {
		conn, err := l.Accept()
		if err != nil {
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	rd := make([]byte, 16*1024)
	for {
		n, err := conn.Read(rd)
		if err != nil {
			return
		}
		fmt.Printf("%s\n", rd[:n])
	}
}
