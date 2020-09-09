package main

import (
	"fmt"
)

func main() {
	var arr1 [10]int

	// 数组循环赋值
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i * i
	}

	// 打印数组信息
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("key=%d,value=%d\n", i, arr1[i])
	}

	// for-range形式的循环
	for k, v := range arr1 {
		fmt.Printf("key=%d,value=%d\n", k, v)
	}
}
