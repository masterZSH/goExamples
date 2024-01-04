package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:6666")
	if err != nil {
		log.Fatal(err)
	}
	rq, err := http.NewRequest("GET", "http://help.dahantc.com/docs/oss/oss-1b9970d8v4ild", nil)
	if err != nil {
		log.Fatal(err)
	}
	err = rq.Write(conn)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			if err != nil {
				continue
			}
			fmt.Printf("%s\n", buf[:n])
		}
	}()

	for {
	}
}
