package main

import (
	"fmt"
	"reflect"
)

type X struct {
	name string
}

func main() {
	a := (*X)(nil)
	fmt.Printf("%+v", reflect.TypeOf(nil))
	fmt.Printf("%+v", reflect.TypeOf(a))
	fmt.Printf("%+v", reflect.ValueOf(a))

}
