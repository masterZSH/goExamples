package main

import (
	"fmt"
)

// Stest struct
type Stest struct {
	f float64
	string
	int
}

// 匿名
type a struct {
	i, j int
}

type b struct {
	a
	k int
}

func main() {
	var sT = &Stest{1.1, "2", 3}
	fmt.Printf("%+v\n", sT)

	var s2 b
	s2.k = 2
	fmt.Printf("%+v\n", s2)
}
