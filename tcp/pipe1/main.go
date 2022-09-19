package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

var key []byte

func init() {
	u := strings.ReplaceAll(uuid.New().String(), "-", "")
	key = []byte(u)
}

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
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// If the key is unique for each ciphertext, then it's ok to use a zero
	// IV.
	var iv [aes.BlockSize]byte
	stream := cipher.NewOFB(block, iv[:])
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn, addr, stream)
	}

}

func handleConn(conn net.Conn, addr string, stream cipher.Stream) {
	tk := time.NewTicker(time.Second)

	// read
	go func() {
		bf := make([]byte, 4*1024)
		for {
			n, err := conn.Read(bf)
			if err != nil {
				return
			}
			bReader := bytes.NewReader(bf[:n])
			reader := &cipher.StreamReader{S: stream, R: bReader}
			// Copy the input to the output stream, decrypting as we go.
			// rd := make([]byte, 1024)
			// n, err = reader.Read(rd)
			// if err != nil {
			// 	panic(err)
			// }
			// fmt.Printf("read %s\n", rd[:n])

			if _, err := io.Copy(os.Stdout, reader); err != nil {
				panic(err)
			}

			// fmt.Printf("receive:%s\n", bf[:n])
		}
	}()
	// mReader := bytes.NewReader([]byte{})

	var bf bytes.Buffer

	writer := &cipher.StreamWriter{S: stream, W: conn}

	for {
		select {
		case _ = <-tk.C:

			// conn.Write([]byte(addr))
			// Encrypt plaintext
			// conn.Write([]byte("123"))
			// conn.Write([]byte("123"))
			n, err := bf.Write([]byte("7"))
			if err != nil {
				continue
			}
			fmt.Printf("write data len %d \n", n)
			fmt.Printf("write %s\n", bf.String())
			io.Copy(writer, &bf)
		}
	}

}

func pipe(src, dst net.Conn) {
	buff := make([]byte, 0xffff)
	// buff := make([]byte, 30*1024)
	for {
		n, err := src.Read(buff)
		if err != nil {
			log.Fatal(err)
		}
		n, err = dst.Write(buff[:n])
		if err != nil {
			log.Fatal(err)
		}
		_ = n
		// fmt.Printf("read %s\n", buff[:n])
	}

}
