package main

import (
	"fmt"
)

func main() {
	var arr1 [5]int
	arr2 := f(arr1) // 不会修改原有的arr1的值，在f函数中做了一次拷贝。修改的是拷贝的值
	fmt.Print(arr1) // [0 0 0 0 0]
	arr3 := fp(&arr1)
	fmt.Print(arr1)          // [0 5 0 0 0]
	fmt.Print(arr1 == *arr2) // false 两个不是一个
	fmt.Print(arr1 == *arr3) // true 两个是一个
	fmt.Print(arr2)          // &[0 5 5 0 0]
}

// 会进行拷贝
func f(a [5]int) *[5]int {
	a[1] = 5
	defer func() {
		a[2] = 5
	}()
	return &a
}

// 通过指针修改
func fp(a *[5]int) *[5]int {
	a[1] = 5
	return a
}
