package main

import "time"

import "fmt"

func main() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	// ch := make(chan int)
	for {
		select {
		case u := <-ticker.C:
			fmt.Println(u)
		}
	}
}
