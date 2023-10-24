package main

import (
	"crypto/tls"
	"fmt"
)

func main() {
	conn, err := tls.Dial("tcp", "localhost:12345", &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		panic("failed to connect: " + err.Error())
	}
	defer conn.Close()

	n, err := conn.Write([]byte("hi"))
	fmt.Print(n, err)
}
