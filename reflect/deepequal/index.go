package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := (*HandleFunc)(nil)
	b := (*HandleFunc)(nil)
	fmt.Print(reflect.DeepEqual(a, b))
}

type HandleFunc func()
