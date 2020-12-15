package main

import (
	"crypto/tls"
	"log"
	"time"
)

func main() {
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}
	conn, err := tls.Dial("tcp", ":6366", conf)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn.Write([]byte("123"))
		time.Sleep(3 * time.Second)
	}
}
