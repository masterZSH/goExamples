package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i interface{}
	i = 1
	fmt.Println(reflect.TypeOf(i))  // int
	fmt.Println(reflect.ValueOf(i)) // 1
	i = "123"
	fmt.Println(reflect.TypeOf(i))  // string
	fmt.Println(reflect.ValueOf(i)) // 123

	// TypeOf 直接给到了我们想要的type类型，如float64、int、各种pointer、struct 等等真实的类型

	// ValueOf 直接给到了我们想要的具体的值

}
