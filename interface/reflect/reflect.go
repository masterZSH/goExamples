// blog: Laws of Reflection
package main

import (
	"fmt"
	"reflect"
)

type Typef struct {
	s1, s2, s3 string
}

func (t *Typef) String() string {
	return t.s1 + " - " + t.s2 + " - " + t.s3
}

var s = Typef{"1", "2", "3"}

func main() {
	// var x float64 = 3.4
	// fmt.Println("type:", reflect.TypeOf(x))
	// v := reflect.ValueOf(&x)
	// if v.CanSet() {
	// 	v.SetFloat(2.222)
	// }

	// fmt.Println(v.Elem())
	// fmt.Println(x)
	KindExample()
	SetField()
}

// SStr declare struct includes three int variables i1,i2,i3
type SStr struct {
	i1, i2, i3 string
}

// Kind的特定类型定义
// const (
// 	Invalid Kind = iota
// 	Bool
// 	Int
// 	Int8
// 	Int16
// 	Int32
// 	Int64
// 	Uint
// 	Uint8
// 	Uint16
// 	Uint32
// 	Uint64
// 	Uintptr
// 	Float32
// 	Float64
// 	Complex64
// 	Complex128
// 	Array
// 	Chan
// 	Func
// 	Interface
// 	Map
// 	Ptr
// 	Slice
// 	String
// 	Struct
// 	UnsafePointer
// )

// KindExample reflect Kind function
func KindExample() {
	var variable = SStr{"1", "2", "3"}
	var rv = reflect.ValueOf(variable)
	var rt = reflect.TypeOf(variable)
	var kind = rv.Kind()
	fmt.Println(rv)   // 输出值 output {1 2 3}
	fmt.Println(rt)   // 输出类型 output main.SInt
	fmt.Println(kind) // 输出特定类型 output struct
}

// NumField NumField返回结构体的字段数量
func NumField() {
	var variable = SStr{"1", "2", "3"}
	var rv = reflect.ValueOf(variable)
	for i := 0; i < rv.NumField(); i++ {
		fmt.Printf("Field %d:%v\n", i, rv.Field(i))
	}
}

//SExport 首字母大写字段可导出 反射Field可以取到
type SExport struct {
	Name string
	Num  int
}

// SetField SStr field
func SetField() {
	var variable = SExport{"foo", 233}
	// 通过反射直接修改原变量variable值
	var r = reflect.ValueOf(&variable).Elem()
	typeOfT := r.Type()
	for i := 0; i < r.NumField(); i++ {
		f := r.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
	// 判断是否可以设置
	if r.Field(0).CanSet() {
		r.Field(0).SetString("123")
	}
	fmt.Printf("%+v", variable)
}

