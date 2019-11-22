package main

import (
	"fmt"
)

type Any interface{}

type Person struct {
	name string
	age  int
}

func main() {
	var v Any
	v = 1
	fmt.Println(v)
	v = "123"
	fmt.Println(v)
	switch t := v.(type) {
	case int:
		fmt.Printf("Type int %T\n", t)
	case string:
		fmt.Printf("Type string %T\n", t)
	case bool:
		fmt.Printf("Type boolean %T\n", t)
	case *Person:
		fmt.Printf("Type pointer to Person %T\n", t)
	default:
		fmt.Printf("Unexpected type %T", t)
	}
}
