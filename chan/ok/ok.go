package main

import "fmt"

func main() {
	ch := make(chan int)
	go input(ch)
	for output := range ch {
		fmt.Print(output)
	}
}

func input(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
