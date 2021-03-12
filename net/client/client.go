package main

import (
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("conn server[%s] success\n", conn.RemoteAddr())

	go readConn(conn)
	go writeConn(conn)
	go heartBeat(conn)

	select {}
}

func readConn(conn net.Conn) {

}

func writeConn(conn net.Conn) {
	conn.Write([]byte("test"))
}

func heartBeat(conn net.Conn) {
	tk := time.NewTicker(5 * time.Second)
	for {
		select {
		case _ = <-tk.C:
			log.Printf("send ping\n")
			conn.Write([]byte("ping"))
		}
	}

}
