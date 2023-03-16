package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	// t 如果是pointer (*a)[low : high : max]
	t := a[1:3:5]
	// a[low : high : max]
	// capacity by setting it to max - low
	// 容量 = max - low

	// []int{2,3}  2       4
	fmt.Println(t, len(t), cap(t))
	fmt.Println(a[1:2])
}
