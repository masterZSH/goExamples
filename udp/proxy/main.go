package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	localAddr := "127.0.0.1:3333"
	targetAddr := "127.0.0.1:5555"
	serve(localAddr, targetAddr)
}

func serve(l, r string) {
	dst, err := net.ResolveUDPAddr("udp", l)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.ListenUDP("udp", dst)
	if err != nil {
		log.Fatal(err)
	}
	rc, err := net.Dial("udp", r)
	if err != nil {
		log.Fatal(err)
	}
	go pipe(rc, conn)
	pipe(conn, rc)
}

func pipe(dst io.Writer, src io.Reader) (err error) {
	buf := make([]byte, 4096)
	for {
		rn, err := src.Read(buf)
		if err != nil {
			break
		}
		fmt.Printf("read %s\n", buf[:rn])
		wn, err := dst.Write(buf[:rn])
		if err != nil {
			break
		}
		fmt.Printf("read %d write %d\n", rn, wn)
	}
	return
}

func handleConn(conn net.UDPConn) {
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {

			continue
		}
		fmt.Printf("read %s\n", buf[:n])
	}

}
