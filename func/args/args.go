package main

import (
	"fmt"
	"reflect"
)

func main() {
	testFunc(1)
}

// 可变参数
func testFunc(arr ...int) {
	fmt.Print(reflect.TypeOf(arr))
}
