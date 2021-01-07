package server

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"
)

func Server() {
	cert, err := tls.LoadX509KeyPair("cert/server.pem", "cert/key.pem")
	if err != nil {
		log.Fatal(err)
	}
	cfg := &tls.Config{Certificates: []tls.Certificate{cert}}
	listener, err := tls.Listen("tcp", ":2666", cfg)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	ct := make([]byte, 0xffffff)
	for {
		n, err := conn.Read(ct)
		if err != nil {
			continue
		}
		rd := bufio.NewReader(bytes.NewReader(ct[:n]))
		r, err := http.ReadRequest(rd)
		fmt.Print(err)
		// not a http request
		if err != nil {
			continue
		}
		c := http.Client{}
		resp, err := c.Do(r)
		fmt.Print(err)
		respD := bufio.NewReader(resp.Body)
		// io.Copy(conn, respD)
		fmt.Print(respD)
	}

}
