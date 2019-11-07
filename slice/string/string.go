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
	fmt.Print(findFileNumbers("../../README.md"))

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
