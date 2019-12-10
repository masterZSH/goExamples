package main

import "fmt"

func main() {
	ch := make(chan int)
	go sum([]int{1, 2, 3, 4, 5}, ch)
	sum := <-ch
	fmt.Print(sum)
}

func sum(bigArray []int, ch chan int) {
	if len(bigArray) == 0 {
		ch <- 0
	} else {
		var sum int
		for _, v := range bigArray {
			sum += v
		}
		ch <- sum
	}
}
