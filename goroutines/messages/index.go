package main

import (
	"fmt"
	"time"
)

//  两个goroutine通过两个无缓冲通道互相
//  转发消息，这个程序能维持每秒多少通信
var (
	send    = make(chan string)
	receive = make(chan string)
)

func main() {

	go func() {
		var i int
		for {
			send <- ""
			m := <-receive
			fmt.Printf("----receive %d %s\n", i, m)
			i++
		}
	}()

	go func() {
		var i int
		for m := range send {
			fmt.Printf("receive %d %s\n", i, m)
			i++
			receive <- ""
		}
	}()

	// 测试一秒通信数量
	time.Sleep(time.Second)

	// 默认情况5162
}
