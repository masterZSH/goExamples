package main

import (
	"fmt"
	"reflect"
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

	// bts := AppendByte([]byte{'a', 'b'}, []byte{'c', 'd'}...)
	// fmt.Print(bts)
	// sl := Filter([]int{1, 2, 3, 4}, isOdd)
	// fmt.Print(sl)
	// var mergeSlice = InsertStringSlice([]int{1, 2}, []int{5, 5, 5, 5, 5}, 3)
	// fmt.Print(mergeSlice)
	// var removeSlice = RemoveStringSlice([]string{"a", "b", "c", "d", "e"}, 1, 2)
	// fmt.Print(removeSlice)

}

func isOdd(i int) bool {
	return i%2 == 0
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

// 用顺序函数过滤容器：s 是前 10 个整型的切片。
// 构造一个函数 Filter，第一个参数是 s，第二个参数是一个 fn func(int) bool，返回满足函数 fn 的元素切片。
// 通过 fn 测试方法测试当整型值是偶数时的情况。

func Filter(s []int, fn func(int) bool) (ns []int) {
	ns = make([]int, 0, len(s))
	for _, v := range s {
		if fn(v) {
			ns = append(ns, v)
		}
	}
	return ns
}

type slice = []int

// 写一个函数 InsertStringSlice 将切片插入到另一个切片的指定位置。
func InsertStringSlice(source slice, target slice, index int) slice {
	var sliceType = reflect.TypeOf(source)
	fmt.Print(sliceType)
	if len(target)-1 < index {
		return target
	}
	leftSlice := target[:index]
	rightSlice := target[index:]
	newSlice := make(slice, len(leftSlice))
	copy(newSlice, leftSlice)
	newSlice = append(newSlice, source...)
	newSlice = append(newSlice, rightSlice...)
	// 扩容
	return newSlice
}

// 写一个函数 RemoveStringSlice 将从 start 到 end 索引的元素从切片 中移除。

//
func RemoveStringSlice(s []string, start int, end int) (newSlice []string) {
	newSlice = make([]string, len(s[:start]))
	copy(newSlice, s[:start])
	defer func() {
		newSlice = append(newSlice, s[end:]...)
	}()
	return
}
