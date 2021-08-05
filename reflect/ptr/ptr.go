package main

import (
	"fmt"
	"reflect"
)

type Nested struct {
	String  string   `selector:"div > p"`
	Classes []string `selector:"li" attr:"class"`
	Struct  *Nested  `selector:"div > div"`
}

func main() {
	n1 := new(Nested)
	n1.String = "123"
	mPtr(n1)

	n2 := Nested{}
	mPtr(n2)

}

func mPtr(v interface{}) {
	rv := reflect.ValueOf(v)
	kd := rv.Kind()
	fmt.Println(kd == reflect.Ptr)
	fmt.Printf("%+v\n", rv)
	// var st reflect.Type
	if kd != reflect.Ptr {
		rv = reflect.ValueOf(&v)
		// st = reflect.TypeOf(&v).Elem()
	}
	// else {
	// 	st = reflect.TypeOf(v).Elem()
	// }

	// sv := rv.Elem()
	// fmt.Printf("%+v\n", sv)
	// for i := 0; i < sv.NumField(); i++ {
	// 	attrV := sv.Field(i)
	// 	if !attrV.CanAddr() || !attrV.CanSet() {
	// 		continue
	// 	}
	// 	fmt.Println(st.Field(i))
	// }

	// st := reflect.TypeOf(v).Elem()

	// fmt.Println(st)
}
