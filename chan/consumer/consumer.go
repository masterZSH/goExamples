package main

import "fmt"

// 用这种习惯用法写一个程序，有两个协程，第一个提供数字 0，10，20，...90
// 并将他们放入通道，第二个协程从通道中读取并打印。
// main() 等待两个协程完成后再结束。
func main() {
	ch := make(chan int)
	done := make(chan bool)
	go Put(ch)
	go Get(ch, done)
	<-done
}

func Put(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i * 10
	}
	// 关闭通道 避免deadlock
	close(ch)
}

func Get(c <-chan int, done chan<- bool) {
	for v := range c {
		fmt.Println(v)
	}
	done <- true
}
