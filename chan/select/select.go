package main

import (
	"fmt"
	"os"
)

func tel(ch chan<- int, quit chan<- bool) {
	for i := 0; i < 15; i++ {
		ch <- i
	}
	quit <- true
}

// select
func main() {
	ch := make(chan int)
	quit := make(chan bool)
	go tel(ch, quit)
	for {
		select {
		case v := <-ch:
			fmt.Print(v)
		case v := <-quit:
			if v == true {
				os.Exit(0)
			}
		}
	}

}
