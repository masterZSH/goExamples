package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {
	var s1 = []int{1, 2, 3, 4, 5}
	s2 := s1
	s2[0] = 10
	fmt.Println(s1, s2)

	s3 := []int{1, 2, 3, 4, 5}
	s4 := slices.Clone(s3)
	s4[0] = 10
	fmt.Println(s3, s4)
}
