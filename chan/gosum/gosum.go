package main

import "fmt"

func main() {
	ch := make(chan int)
	go sum(1, 2, ch)
	sum := <-ch
	fmt.Print(sum)
}

func sum(i, j int, ch chan int) {
	sum := i + j
	ch <- sum
}
