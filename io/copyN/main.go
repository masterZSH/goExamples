package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	rd := bytes.NewReader([]byte("1234"))
	n, err := io.CopyN(os.Stdout, rd, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("written number %d\n", n)
}
