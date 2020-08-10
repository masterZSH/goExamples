package main

import "fmt"

// slice是一种轻量级的数据结构，可以用来访问数组
// 的部分或者全部的元素，而这个数组成为slice的底层数组
// slice有三个属性
// 指针 指向底层数组
// 长度 len函数
// 容量 cap函数
//

func main() {
	var slice = []string{"a", "b", "c", "d"}

	// slice1 创建一个新的切片
	slice1 := slice[:2] // [a,b]

	slice1 = append(slice1, "e")
	fmt.Print(slice1) // [a,b,e]
	fmt.Print(slice)
}
