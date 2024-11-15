package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := (*HandleFunc)(nil)
	b := (*HandleFunc)(nil)
	fmt.Print(reflect.DeepEqual(a, b))

	i1 := []int{1, 2, 3}
	i2 := []int{1, 2, 3}
	fmt.Println(reflect.DeepEqual(i1, i2))
}

type HandleFunc func()
