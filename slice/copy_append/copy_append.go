package main

import (
	"fmt"
)

func main() {
	sl_from := []int{1, 2, 3}
	sl_to := make([]int, 10)
	// 将sl_from 的元素赋值到sl_to
	n := copy(sl_to, sl_from)
	fmt.Println(sl_to)
	// [1,2,3,0,0,0,0,0,0,0]
	fmt.Printf("Copied %d elements\n", n) // n == 3

	// 追加
	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)
	// [1,2,3,4,5,6]

	// func append(s[]T, x ...T) []T
	// 其中 append 方法将 0 个或多个具有相同类型
	// s 的元素追加到切片后面并且返回新的切片；
	// 追加的元素必须和原切片的元素同类型。
	// 如果 s 的容量不足以存储新增元素，
	// append 会分配新的切片来保证已有切片元素和新增元素的存储。
	// 因此，返回的切片可能已经指向一个不同的相关数组了。
	// append 方法总是返回成功，除非系统内存耗尽了。

	bts := AppendByte([]byte{'a', 'b'}, []byte{'c', 'd'}...)
	fmt.Print(bts)
}

func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]byte, (n)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}

// 给定一个slices []int 和一个 int 类型的因子，
// 扩展 s 使其长度为 len(s) * factor。
func ap(slice []int, fac int) []int {
	//
	newSlice := make([]int, len(slice)*fac)
	copy(newSlice, slice)
	slice = newSlice
	return slice
}
