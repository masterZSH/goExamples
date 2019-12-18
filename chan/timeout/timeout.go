package main

import (
	"fmt"
	"time"
)

// 1s后向timeout channel加入true
func f(timeout chan bool) {
	time.Sleep(time.Second)
	timeout <- true
}

// s 2s后向ch写入1
func s(ch chan int) {
	time.Sleep(2 * time.Second)
	ch <- 1
}

func main() {
	timeout := make(chan bool, 1)
	ch := make(chan int, 1)
	go f(timeout)
	go s(ch)
	select {
	case i := <-ch:
		fmt.Print(i)
	case <-timeout:
		fmt.Print("timeout")
	}
}
