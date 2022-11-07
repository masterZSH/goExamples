package main

import (
	"fmt"
	"time"
)

type Empty interface{}

var (
	empty Empty
)

const N int = 100

func main() {
	data := make([]float64, N)
	res := make([]float64, N)
	sem := make(chan Empty, N)
	data = []float64{1.1, 2.2}
	for i, xi := range data {
		go func(i int, xi float64) {
			res[i] = doSomething(i, xi)
			sem <- empty
		}(i, xi)
	}
	fmt.Print(data)
}

func doSomething(k int, v float64) float64 {
	time.Sleep(5e8)
	return v * 2
}
