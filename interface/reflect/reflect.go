// blog: Laws of Reflection
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	v := reflect.ValueOf(&x)
	if v.CanSet() {
		v.SetFloat(2.222)
	}

	fmt.Println(v.Elem())
}
