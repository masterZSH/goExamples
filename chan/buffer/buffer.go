package main

// 写一个通道证明它的阻塞性，开启一个协程接收通道的数据，持续 15 秒，
// 然后给通道放入一个值。在不同的阶段打印消息并观察输出。
import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 100)
	go getData(ch)
	go setData(ch)
	time.Sleep(1e9)
}

func setData(ch chan string) {
	ch <- "123"
}

func getData(ch chan string) {
	fmt.Print(<-ch)
	time.Sleep(19e9)
}
