package main

import (
	"log"
)

func err() {
	panic("err")
}

func main() {
	defer func() {
		if e := recover(); e != nil {
			log.Printf("panic: %v", e)
		}
	}()
	err()
}
