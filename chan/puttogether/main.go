package main

import (
	"fmt"
	"time"
)

func ToChansTimedTimerSelect[T any](d time.Duration, message T, a, b chan T) (written int) {
	t := time.NewTimer(d)
	for i := 0; i < 2; i++ {
		select {
		case a <- message:
			a = nil
		case b <- message:
			b = nil
		case <-t.C:
			return i
		}
	}
	t.Stop()
	return 2
}

func main() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	n := ToChansTimedTimerSelect[string](
		time.Second,
		"foo",
		ch1, ch2,
	)

	fmt.Println(n)
	v1, ok := <-ch1
	fmt.Println(v1, ok)
	v2, ok := <-ch2
	fmt.Println(v2, ok)
}
