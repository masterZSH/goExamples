package main

import (
	"fmt"
	"time"
)

func main() {
	cancel := make(chan bool)
	// 开启协程
	go r(cancel)

	time.Sleep(10 * time.Second)
	// 写入关闭通道true
	cancel <- true
	time.Sleep(time.Second)
}

func r(ch chan bool) {
	for {
		select {
		case <-ch:
			fmt.Println("cancel\n")
			return
		default:
			fmt.Println("start")
			time.Sleep(2 * time.Second)
			fmt.Println("end")
		}
	}
}
