package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

// TestByte 10 byte
//go:notinheap
type TestByte [10]byte

func main() {

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	for {
		var t TestByte
		t[0] = 1
	}
}
