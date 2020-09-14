package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8989")
	delErr(err)
	bs := new(bytes.Buffer)
	content := []byte("123")
	err = binary.Write(bs, binary.LittleEndian, uint32(len(content)))
	delErr(err)
	msg := append(bs.Bytes(), content...)
	l, err := conn.Write(msg)
	delErr(err)
	fmt.Printf("write len: %d", l)
}

func delErr(err error) {
	if err != nil {
		log.Fatalf("err:%+v\n", err)
	}
}
