package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i int = 1
	t := reflect.TypeOf(i)
	fmt.Println(t)

	v := reflect.ValueOf(i)
	fmt.Println(v)

	fmt.Println(v.CanConvert(reflect.TypeOf(bool(true))))
}
