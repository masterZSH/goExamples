package main

import (
	"fmt"
)

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func Sort(data Sorter) {
	for pass := 1; pass < data.Len(); pass++ {
		for i := 0; i < data.Len()-pass; i++ {
			if data.Less(i+1, i) {
				data.Swap(i, i+1)
			}
		}
	}
}

type L []int

func (l L) Len() int {
	return len(l)
}

func (l L) Less(i, j int) bool {
	return i < j
}

func (l L) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func main() {
	var i Sorter = L([]int{3, 2, 1, 5, 4})
	Sort(i)
	fmt.Print(i)
}
