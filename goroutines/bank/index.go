package main

import "fmt"

var (
	balance int
	sign    = make(chan struct{}, 1)
)

func main() {
	go deposit(1)
	b1 := ba()
	fmt.Println(b1)
}

func deposit(i int) {
	sign <- struct{}{}
	balance = balance + i
	<-sign
}

func ba() int {
	sign <- struct{}{}
	b := balance
	<-sign
	return b
}
