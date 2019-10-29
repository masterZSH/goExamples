package main

import "fmt"

// fobinacci_funcarray.go: 为练习 7.3 写一个新的版本，
// 主函数调用一个使用序列个数作为参数的函数，
// 该函数返回一个大小为序列个数的 Fibonacci 切片。
func returnFib(size int) (r []int) {
	if size < 3 {
		return []int{1, 1}
	}
	size = size - 1
	a := returnFib(size)[:]
	b := append(a, a[len(a)-1]+a[len(a)-2])
	return b
}

func main() {
	// fmt.Print(",")
	fmt.Print(returnFib(5))
}
