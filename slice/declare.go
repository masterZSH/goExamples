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

	// var s1 []int = arr1[:2]
	// var s2 []int = arr1[2:]

	// arr1[:2] + arr1[2:] // i是一个整数且: 0 <= i <= len(arr1)
	// len(arr1) <= cap(arr1)

	// 切片在内存中的组织方式实际上是一个有 3 个域的结构体：
	// 指向相关数组的指针，切片长度以及切片容量。
	// 下图给出了一个长度为 2，容量为 4 的切片y。
	fmt.Println(len(slice2))

	// 不要用指针指向切片，切片本身就是引用类型
	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	fmt.Println(b[1:4], b[:2], b[2:], b[:])
	// b[1:4] ['o','l','a']
	// b[:2] ['g','o']
	// b[2:] ['l','a','n','g']
	// b[:]  ['g','o','l','a','n','g']

	// 使用make定义切片
	var makeSlice1 = make([]int, 10, 20)
	fmt.Println(makeSlice1)
	// 使用new定义切片
	var makeSlice2 = new([20]int)[0:10]
	fmt.Println(makeSlice2)

	var makeSlice3 = make([]byte, 5)
	makeSlice3 = makeSlice3[2:4]
	fmt.Println(len(makeSlice3), cap(makeSlice3))

}
