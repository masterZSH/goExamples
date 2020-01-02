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
	// 注意缓冲大小设置为 1 是必要的，
	// 可以避免协程死锁以及确保超时的通道可以被垃圾回收
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
