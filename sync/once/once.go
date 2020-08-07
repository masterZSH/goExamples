package main

import (
	"fmt"
	"sync"
)

var one sync.Once

type T struct {
	name string
	age  int
}

var t *T

func main() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("只执行一次")
	}
	done := make(chan bool)
	// 10个协程
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}

	// 10次执行完了
	for i := 0; i < 10; i++ {
		<-done
	}
}

func getT() *T {
	one.Do(func() {
		t = new(T)
	})
	return t
}
