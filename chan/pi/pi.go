package main

import (
	"fmt"
	"math"
)

// pi计算 随着计算越来越精确
func pi(ch chan float64) {
	var r float64
	for k := 0; ; k++ {
		i := float64(k)
		item := 4 * (math.Pow(-1, i) / (2.0*i + 1.0))
		r += float64(item)
		ch <- r
	}
}

func getData(ch chan float64) {
	for i := range ch {
		fmt.Printf("pi：%f\n", i)
	}
}

func main() {
	ch := make(chan float64)
	go pi(ch)
	getData(ch)
}
