package main

import "time"

import "fmt"

// MAXNUM 最大
const MAXNUM = 50

var sem = make(chan int, MAXNUM)

func process() {
	time.Sleep(time.Second)
	fmt.Print("1")
}

func server() {
	for {
		go process()
	}
}
func handle() {
	sem <- 1
	go process()
	<-sem
}

func main() {
	go server()

}
