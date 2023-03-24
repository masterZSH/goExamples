package main

import (
	"fmt"
	"sync"
)

func firstComeFirstServedSelect[T any](message T, a, b chan<- T) {
	for i := 0; i < 2; i++ {
		select {
		case a <- message:
			a = nil
		case b <- message:
			b = nil
		}
	}
}

func firstComeFirstServedGoroutines[T any](message T, a, b chan<- T) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() { a <- message; wg.Done() }()
	go func() { b <- message; wg.Done() }()
	wg.Wait()
}

func main() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)
	firstComeFirstServedSelect[string](
		"123",
		c1,
		c2,
	)
	v1, ok := <-c1
	fmt.Println("c1", v1, ok)
	v2, ok := <-c2
	fmt.Println("c2", v2, ok)

	c3 := make(chan string, 1)
	c4 := make(chan string, 1)
	firstComeFirstServedGoroutines[string]("456", c3, c4)

	v3, ok := <-c3
	fmt.Println("c3", v3, ok)
	v4, ok := <-c4
	fmt.Println("c4", v4, ok)

}
