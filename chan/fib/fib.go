package main

import "fmt"

func fib(n int) (v, p int) {
	if n <= 1 {
		v = 1
	} else {
		v1, _ := fib(n - 1)
		v2, _ := fib(n - 2)
		v = v1 + v2
	}
	p = n
	return
}

func fibT(s int, ch chan int) {
	for i := 0; i <= s; i++ {
		v, _ := fib(i)
		ch <- v
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go fibT(10, ch)
	for i := range ch {
		fmt.Print(i)
	}
}
