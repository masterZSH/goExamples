package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("start:\n")
	// 无缓冲同步
	ch := make(chan string)
	go t(ch)
	s := <-ch
	fmt.Printf("end %s<-ch\n", s)
}

func t(ch chan string) {
	fmt.Printf("ch <- \"test\"\n")
	time.Sleep(1 * time.Second)
	ch <- "test"
}
