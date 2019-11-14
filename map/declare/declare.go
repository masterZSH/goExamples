package main

import (
	"fmt"
)

func main() {
	// map是引用类型
	// var map1 map[string]int
	mapM := make(map[string]int)
	mapM["1"] = 1
	fmt.Print(mapM)
	// 	map 可以根据新增的 key-value 对动态的伸缩
	// 可以分配内容初始化容量
	map2d := make(map[int][]int, 1)
	map2d[0] = []int{1, 2, 3}
	fmt.Print(map2d)
}
