package main

import (
	"fmt"
	"reflect"
)

// 为 int 和 string 构建一个把 int 值加倍和将字符串值与其自身连接
// （译者注：即"abc"变成"abcabc"）的 map 函数 mapFunc。
func mapFunc(arg interface{}) interface{} {
	// var tp = reflect.TypeOf(arg)
	// 判断类型
	// if reflect.String == tp.Kind() {
	// fmt.Printf("%+v", p)
	// }
	// 或者
	// if v, ok := arg.(string); ok {
	// 	fmt.Printf("%+v", v)
	// }
	var rarg = reflect.ValueOf(arg)
	switch arg.(type) {
	case string:
		return rarg.String() + rarg.String()
	case int:
		return rarg.Int() * 2
	default:
		return arg
	}
}

// 稍微改变练习允许 mapFunc 接收不定数量的 items。
func mapFuncArgs(args ...interface{}) ([]interface{}){
	for k,arg := range args{
		var rarg = reflect.ValueOf(arg)
		switch arg.(type) {
		case string:
			args[k] = rarg.String() + rarg.String()
		case int:
			args[k] = rarg.Int() * 2
		}
	}
	return args
}


func main() {
	// fmt.Print(mapFunc(2))     // 4
	// fmt.Print(mapFunc("123")) // 123123

	fmt.Print(mapFuncArgs(2,"123"))
}
