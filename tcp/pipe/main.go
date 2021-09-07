package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	go l(":12666")
	go l(":12777")

	conn1, err := net.Dial("tcp", ":12666")
	if err != nil {
		log.Fatal(err)
	}
	conn2, err := net.Dial("tcp", ":12777")
	if err != nil {
		log.Fatal(err)
	}

	go pipe(conn1, conn2)
	go pipe(conn2, conn1)
	for {
	}
}

func l(addr string) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn, addr)
	}

}

func handleConn(conn net.Conn, addr string) {
	tk := time.NewTicker(time.Second)
	go func() {
		bf := make([]byte, 1024)
		for {
			n, err := conn.Read(bf)
			if err != nil {
				return
			}
			fmt.Printf("receive:%s\n", bf[:n])
		}
	}()

	for {
		select {
		case _ = <-tk.C:
			conn.Write([]byte(addr))
		}
	}

}

func pipe(src, dst net.Conn) {
	buff := make([]byte, 0xffff)
	for {
		n, err := src.Read(buff)
		if err != nil {
			log.Fatal(err)
		}
		n, err = dst.Write(buff[:n])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", buff[:n])
	}

}
