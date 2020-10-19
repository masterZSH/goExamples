package main

import (
	"io"
	"log"
	"os"
)

func main() {
	f1, err := os.Open("1")
	if err != nil {
		log.Fatalf("open file 1 failure")
	}
	f2, err := os.OpenFile("2", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("open file 2 failure")
	}
	// pipe
	num, err := io.Copy(f2, f1)
	if err != nil {
		log.Fatalf("copy file failure")
	}

	log.Printf("copy file success,write byte num %d\n", num)
}
