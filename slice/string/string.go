package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

// 1. 字符串本质上是字节数组 所以可以用切片形式访问
// 2. 字符串一个字符串实际上是一个双字结构，
// 即一个指向实际数据的指针和记录字符串长度的整数。
// 3. 字符串不可修改，可以转为字节数组修改

func main() {
	// var testStr string = "test"
	// fmt.Print(testStr[2:])
	// var byte1 = []byte{'a'}
	// fmt.Println(string(append(byte1, testStr...)))
	// var intArr []int = []int{1, 5, 3, 4, 2}
	// sort.Ints(intArr)
	// fmt.Println(intArr)
	// insertArr := insert([]int{1, 2, 3}, 2, 4)
	// fmt.Print(insertArr)
	// fmt.Print(getLast([]int{1, 2, 3}))
	// fmt.Print(findFileNumbers("../../README.md"))
	// fmt.Print(reverse("Google"))
	// fmt.Print(copy([]byte{'a', 'b', 'c', 'c'}))
	fmt.Print(bubble([]int{3, 2, 1, 5, 4}))
	fmt.Print(mapFunc(mf, []int{1, 2, 3, 4}))
}

// 删除位于索引 index 的元素
func deleteIndex(arr []int, index int) []int {
	arr = append(arr[:index], arr[index+1:]...)
	return arr
}

// 扩展切片
func ext(arr []int, len int) []int {
	return append(arr, make([]int, len)...)
}

// 删除切片中的元素
func del(arr []int, start int, end int) []int {
	arr = append(arr[:start], arr[end:]...)
	return arr
}

func insert(arr []int, i int, v int) []int {
	arr = append(arr[:i], append([]int{v}, arr[i:]...)...)
	return arr
}

func getLast(arr []int) int {
	return arr[len(arr)-1]
}

func findFileNumbers(filename string) [][]byte {
	var reg = regexp.MustCompile("[0-9]")
	bytes, _ := ioutil.ReadFile(filename)
	return reg.FindAll(bytes, len(bytes))
}

// 编写一个函数，要求其接受两个参数，
// 原始字符串 str 和分割索引 i，然后返回两个分割后的字符串。
func split(str string, i int) (s1 string, s2 string) {
	return str[:i], str[i:]
}

// 假设有字符串 str，
// 那么 str[len(str)/2:] + str[:len(str)/2]
// 的结果是什么？
func compact() {
	var str string = "test1"
	// st1te
	fmt.Print(str[len(str)/2:] + str[:len(str)/2])
}

// 编写一个程序，要求能够反转字符串，
// 即将 “Google” 转换成 “elgooG”（提示：使用 []byte 类型的切片）。
// 如果您使用两个切片来实现反转，
//请再尝试使用一个切片（提示：使用交换法）。
// 如果您想要反转 Unicode 编码的字符串，
//请使用 []int32 类型的切片。

func reverse(str string) string {
	slice := []byte(str)
	for k := 0; k < len(str)/2; k++ {
		// 交换两个数字
		temp := slice[k]
		slice[k] = slice[(len(str)-1)-k]
		slice[(len(str)-1)-k] = temp
	}
	return string(slice)
}

// 编写一个程序，要求能够遍历一个字符数组，
// 并将当前字符和前一个字符不相同的字符拷贝至另一个数组。
func copy(bs []byte) []byte {
	var newBytes = []byte{}
	for k := 0; k < len(bs); k++ {
		if k != 0 && (bs[k] != bs[k-1]) {
			newBytes = append(newBytes, bs[k])
		}
	}
	return newBytes
}

// 编写一个程序，使用冒泡排序的方法排序一个包含整数的切片
// （算法的定义可参考 维基百科）。

func bubble(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr)-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				// swap
				// var temp = arr[j]
				// arr[j] = arr[j+1]
				// arr[j+1] = temp
			}
		}
	}
	return arr
}

// 在函数式编程语言中，
// 一个 map-function 是指能够接受一个函数原型和一个列表，
// 并使用列表中的值依次执行函数原型，
// 公式为：map ( F(), (e1,e2, . . . ,en) ) = ( F(e1), F(e2), ... F(en) )。

func mapFunc(mapf func(int) int, arr []int) []int {
	var slice = make([]int, len(arr))
	for k, v := range arr {
		slice[k] = mapf(v)
	}
	return slice
}

func mf(i int) int {
	return i * 10
}
