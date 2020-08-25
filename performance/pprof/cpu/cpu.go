package main

import (
	"fmt"

	"github.com/pkg/profile"
)

func main() {
	defer profile.Start().Stop()

	sl := makeSlice()
	fmt.Printf("sum = %d\n", sumSlice(sl))
}

func makeSlice() []int {
	sl := make([]int, 10000000)
	for idx := range sl {
		sl[idx] = idx
	}
	return sl
}

func sumSlice(sl []int) int {
	sum := 0
	for _, x := range sl {
		sum += x
	}
	return sum
}
