package main

import "fmt"

func main() {
	var arr = [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	s1 := arr[1:]
	fmt.Println(s1)
	s1[2] = 999
	s1 = append(s1, []int{11, 12}...)
	fmt.Println(s1)
	// 扩容了底层数组已经变了
	s1[2] = 888
	fmt.Println(s1)
	fmt.Print(arr)
}
