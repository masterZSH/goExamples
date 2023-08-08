package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Foo[T constraints.Integer] struct {
	Data T
}

func main() {
	f1 := Foo[int]{int(1)}
	f2 := Foo[int32]{int32(1)}

	fmt.Println(f1, f2)
}
