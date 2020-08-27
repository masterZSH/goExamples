package main

import (
	"fmt"
	"reflect"
)

type A struct {
	Name string
}

func main() {
	foo := []A{
		{
			Name: "123",
		},
		{
			Name: "456",
		},
	}

	bar := []A{
		{
			Name: "789",
		},
		{
			Name: "000",
		},
	}

	ty1 := reflect.ValueOf(foo)
	ty2 := reflect.ValueOf(bar)

	// 复制前
	fmt.Printf("复制前----------\n")
	fmt.Print(ty2)

	num := reflect.Copy(ty2, ty1)
	fmt.Printf("\n复制数量:%d", num)

	// 复制后
	fmt.Printf("\n复制后----------\n")
	fmt.Print(ty2)
}
