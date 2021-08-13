package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {

	reader := strings.NewReader("foo bar")
	buf := make([]byte, 4)
	n, err := io.ReadFull(reader, buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf[:n])
	// minimal read size bigger than io.Reader stream
	bigBuf := make([]byte, 1024)

	n, err = io.ReadFull(reader, bigBuf)
	if err != nil {
		//error: unexpected EOF
		log.Fatal(err)
	}

}
