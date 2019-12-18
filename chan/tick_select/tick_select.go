package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1 * time.Second)
	after := time.After(2 * time.Second)
	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-after:
			fmt.Println("after")
		default:
			fmt.Println("    .")
			time.Sleep(5e7)
		}
	}
}
