package main

import (
	"sync"
	"time"
)

const (
	i1 = iota
	i2
	i3
	i4
)

func main() {
	// ch := make(chan int)
	// go run(ch)
	// for i := range ch {
	// 	fmt.Printf("%d ", i)
	// }

	var wg sync.WaitGroup
	wg.Add(delta)

	ch := make(chan int)
	go f(ch)
	<-ch

}

func run(ch chan int) {
	for i := 0; i < 1000; i++ {
		ch <- i
	}
	close(ch)
}

func f(ch chan int) {
	time.Sleep(time.Second)
	ch <- 1
}
