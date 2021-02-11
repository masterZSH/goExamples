package main

import "fmt"

// 并行计算1-n
//
//
// cpu nums cpu数量
const cpus = 3

func main() {
	// 1-100的和
	sum := calcAll(100)
	fmt.Printf("sum:%d\n", sum)
}

// 1 2 3 4 5 6 7 8 9 10

func calcAll(v int) (sum int) {
	c := make(chan int, cpus)
	for i := 0; i < cpus; i++ {
		// 拆分计算
		go calc(i*v/cpus, (i+1)*v/cpus, c)
	}
	for i := 0; i < cpus; i++ {
		sum += <-c
	}
	return
}

func calc(l int, r int, c chan int) {
	var sum int
	for i := l; i < r; i++ {
		sum += i
	}
	c <- sum
}
