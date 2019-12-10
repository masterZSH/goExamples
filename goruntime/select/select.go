package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// 能感觉到明显的性能提升
	// 2->3800 4->7800
	runtime.GOMAXPROCS(4)
	ch1 := make(chan int)
	ch2 := make(chan int)
	go pump1(ch1)
	go pump2(ch2)
	go suck(ch1, ch2)
	time.Sleep(1e9) // 1s
}

func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}

func pump2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 5
	}
}

func suck(ch1, ch2 chan int) {
	for i := 0; ; i++ {
		select {
		case v := <-ch1:
			fmt.Printf("%d - Received on channel 1: %d\n", i, v)
		case v := <-ch2:
			fmt.Printf("%d - Received on channel 2: %d\n", i, v)
		}
	}
}
