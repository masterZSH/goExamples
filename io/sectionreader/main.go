package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	s := io.NewSectionReader(r, 5, 17)

	fmt.Println(s.Size())

	buf := make([]byte, 100)
	if _, err := s.Read(buf); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf)
}
