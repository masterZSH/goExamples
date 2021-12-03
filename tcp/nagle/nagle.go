package nagle

import (
	"fmt"
	"net"
)

func Start() {
	l, err := net.Listen("tcp", ":6666")
	if err != nil {
		panic(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	// Nagle's algorithm
	if v, ok := conn.(*net.TCPConn); ok {
		v.SetNoDelay(true)
	}
	buf := make([]byte, 1024*4)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			return
		}
		if n != 0 {
			fmt.Printf("received: %s\n", buf[:n])
		}
	}
}
