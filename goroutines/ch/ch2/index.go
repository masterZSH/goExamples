package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// 超时通道 超时时候写入该通道
	timeout:=make(chan bool)
	ch:=make(chan int,10)
	// 5s后超时写入超时通道
	go timeoutFunc(timeout)
	// 每隔1s发送数字到ch通道
	go sendNum(ch)
	for{
		select {
		case num:=<-ch:
			fmt.Printf("get num:%d\n",num)
		case <-timeout:
			// 超时 退出程序
			fmt.Printf("timeout\n")
			os.Exit(1)
		}
	}
}

func sendNum(ch chan int){
	var i int 
	for{
		ch<-i
		time.Sleep(time.Second)
		i++
	}
}

func timeoutFunc(timeout chan bool)  {
	time.Sleep(5*time.Second)
	timeout<-true
}