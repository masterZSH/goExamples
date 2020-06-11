package main

import "fmt"

func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	go counter(ch1)
	go sq(ch1, ch2)
	print(ch2)
}

func counter(ch chan int) {
	var i int
	for {
		ch <- i
		i++
	}
}

func sq(in chan int, out chan int) {
	for item := range in {
		out <- item * item
	}
}

func print(in chan int) {
	for item := range in {
		fmt.Printf("num:%d\n", item)
	}
}
