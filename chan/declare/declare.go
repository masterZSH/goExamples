package main

import "fmt"
import "time"

func main() {
	// chanString := make(chan string)
	ch := make(chan string)
	go sendData(ch)
	go getData(ch)
	time.Sleep(1e9)
}

func sendData(ch chan string) {
	ch <- "string1"
	ch <- "string2"
	ch <- "string3"
	ch <- "string4"
}

func getData(ch chan string) {
	var i string
	for {
		i = <-ch
		fmt.Printf("%s", i)
	}
}

// 解释一下为什么如果在函数 getData() 的一开始插入
// time.Sleep(2e9)，不会出现错误但也没有输出呢。

// 此时main方法中的time.Sleep(1e9)是1smain函数结束后会销毁协程
// 所以没有输出
