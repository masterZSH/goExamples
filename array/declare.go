package main

import (
	"fmt"
)

func main() {
	// 申明格式 var identifier [len]type
	var arr1 [5]int
	var arr2 [2]byte
	fmt.Print(arr1)
	fmt.Print(arr2)

	// go中数组是值类型，可以通过new()创建,返回的是指针
	var newArr1 = new([5]int)
	fmt.Print(newArr1)
	newArr2 := *newArr1 // 把newArr1赋值给newArr2时会做一次内存拷贝
	newArr2[1] = 2
	fmt.Print(newArr2) // [0 2 0 0 0]
	fmt.Print(newArr1) // &[0 0 0 0 0]

	// 修改原有的newArr1
	newArr1[1] = 2
	fmt.Print(newArr1) // &[0 2 0 0 0]

	// 数组常量初始化
	var newArr3 = [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(newArr3); i++ {
		fmt.Printf("key at %d is %d\n", i, newArr3[i])
	}

	//从左边开始 如果赋值的数量不足len(arr)，其他的元素会被初始化为0
	var newArr4 = [5]int{1, 2} // [1 2 0 0 0]
	fmt.Print(newArr4)

	// 数组长度...
	var newArr5 = [...]int{1, 2}
	fmt.Print(newArr5)

	// 多维数组
	const (
		WIDTH  = 1920
		HEIGHT = 1080
	)
	type pixel int
	var screen [WIDTH][HEIGHT]pixel
	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WIDTH; j++ {
			screen[j][i] = 0
		}
	}
	fmt.Print(screen)

}
