package main

import "fmt"

// 当结构体的命名以大写字母开头时，该结构体在包外可见。
type struct1 struct {
	Name string
	Age  int
}

func main() {
	m := &struct1{"123", 1}
	//
	fmt.Printf("%+v", m)
}

// 通常情况下，为每个结构体定义一个构建函数，并推荐使用构建函数初始化结构体
func createStruct(name string, age int) *struct1 {
	return &struct1{name, age}
}
