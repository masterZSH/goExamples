package main

import (
	"fmt"
	"log"
	"strings"
	"testing/iotest"
)

func main() {
	r := strings.NewReader("123456")
	hr := iotest.HalfReader(r)
	p := make([]byte, 10)
	// 	return r.r.Read(p[0 : (len(p)+1)/2])
	n, err := hr.Read(p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", p[:n])
}
