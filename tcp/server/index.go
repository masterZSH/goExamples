package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8989")
	delErr(err)
	for {
		conn, err := l.Accept()
		delErr(err)
		go handleConn(conn)
	}
}

func delErr(err error) {
	if err != nil {
		log.Fatalf("err:%+v\n", err)
	}
}

func handleConn(conn net.Conn) {
	// header
	var header []byte
	headerLen := 4
	for {
		tmp := make([]byte, headerLen-len(header))
		n, err := conn.Read(tmp)
		delErr(err)
		header = append(header, tmp[:n]...)
		if headerLen == len(header) {
			break
		}
	}
	var numBytesUint32 uint32
	rbuf := bytes.NewReader(header)
	err := binary.Read(rbuf, binary.LittleEndian, &numBytesUint32)
	delErr(err)
	numBytes := int(numBytesUint32)
	fmt.Printf("header : %v\n", header)
	fmt.Printf("content len : %d\n", numBytes)
	buf := make([]byte, 0)
	for {
		tmp := make([]byte, numBytes-len(buf))
		n, err := conn.Read(tmp)
		delErr(err)
		buf = append(buf, tmp[:n]...)
		if numBytes == len(buf) {
			break
		}
	}
	fmt.Printf("content:%s", buf)
}
