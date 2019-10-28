package main

import (
	"fmt"
)

func main() {
	// 切片（slice）是对数组一个连续片段的引用
	//（该数组我们称之为相关数组，通常是匿名的），
	// 所以切片是一个引用类型
	//（因此更类似于 C/C++ 中的数组类型，或者 Python 中的 list 类型）。
	// 这个片段可以是整个数组，
	// 或者是由起始和终止索引标识的一些项的子集。
	// 需要注意的是，终止索引标识的项不包括在切片内。
	// 切片提供了一个相关数组的动态窗口。
	// 切片是可索引的，并且可以由 len() 函数获取长度。
	// 切片提供了计算容量的函数 cap() 可以测量切片最长可以达到多少：
	// 它等于切片的长度 + 数组除切片之外的长度。

	// 声明一个切片
	var slice1 []int
	fmt.Println(slice1, cap(slice1), len(slice1)) // [] 0 0

	// 从数组中声明
	// 格式 var slice1 []type = arr1[start:end]。 start到end-1索引之间的集合
	var arr1 = [5]int{1, 2, 3, 4, 5}
	var slice2 = arr1[:]         // slice2 等于完整的 arr1数组 另一种写法是 slice1 = &arr1
	var slice3 []int = arr1[3:4] // 下标3的元素集合 [4]
	var slice4 []int = arr1[3:5] // 下标3-4的元素集合[4,5]
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
	// 不可超过切片索引给切片赋值
	// slice4[3] = 1                 // range error 会报range error

}
