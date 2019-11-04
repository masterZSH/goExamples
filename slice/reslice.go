package main

import (
	"fmt"
)

func main() {
	// 切片可以反复扩展直到占据整个相关数组。
	slice1 := make([]int, 0, 10)
	fmt.Print(slice1)
	// 扩充
	// slice1 = slice1[0 : len(slice1)+1]
	// fmt.Print(slice1)
	for i := 0; i < cap(slice1); i++ {
		slice1 = slice1[0 : i+1]
		slice1[i] = i
		fmt.Printf("len of slice is %d\n", len(slice1))
	}

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	var slice2 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var sa = slice2[5:7]
	fmt.Print(sa, len(sa), cap(sa))
	// [6,7] 2 5
	sa = sa[0:4]
	fmt.Print(sa, len(sa), cap(sa))

	// 问题 a是一个切片，那么a[n:n]的长度是多少呢？
	// 猜测为 0,cap为未切片前的长度-n

	// 开始测试
	slice3 := []int{1, 2, 3, 4}
	fmt.Println(slice3, len(slice3), cap(slice3))
	s3 := slice3[1:1]
	fmt.Println(s3, len(s3), cap(s3))
	// [] 0 3

	// a[n:n+1] 的长度又是多少？
	// 猜测为1

	// 开始验证
	s4 := slice3[1:2]
	fmt.Println(s4, len(s4), cap(s4))

}
