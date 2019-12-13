package main

import "fmt"

func sendData(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func getData(ch <-chan int) {
	// for {
	// 	input, open := <-ch
	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Printf("%d ", input)
	// }
	for i := range ch {
		fmt.Printf("%d ", i)
	}
}

func main() {
	ch := make(chan int)
	go sendData(ch)
	getData(ch)
}
