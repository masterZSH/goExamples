package main

import "fmt"

func main() {
	s1 := make([]int, 10)
	s2 := make([]int, 5)
	copy(s2, s1)
	fmt.Print(s2)
}
