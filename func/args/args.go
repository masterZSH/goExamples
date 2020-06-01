package main

import (
	"fmt"
	"reflect"
)

func main() {
	testFunc(1)
	testError()
}

// 可变参数
func testFunc(arr ...int) {
	fmt.Print(reflect.TypeOf(arr))
}

func testError() {
	e := &NeError{"123"}
	fmt.Print(e.Error())
}

type NeError struct {
	message string
}

func (n *NeError) Error() string { return n.message }
