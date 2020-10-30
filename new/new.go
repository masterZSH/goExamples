package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	// allocates zeroed storage for a new item of type T and returns its address
	// a value of type *T
	// zero value
	// 没有初始化内存，只为空
	//  it does not initialize the memory, it only zeros it
	pf := new(os.File)
	fmt.Print(pf)
	fmt.Printf("%p\n", pf)

	// *os.File
	fmt.Println(reflect.TypeOf(pf))

	// new(File) and &File{} are equivalent
	// new(os.File) &os.File{}

	// make会初始化数据
	// it returns an initialized (not zeroed) value of type T (not *T)
	s := make([]int, 10)

	// []int
	fmt.Println(reflect.TypeOf(s))

	// references to data structures that must be initialized before use

	//  It creates slices, maps, and channels only

	//  new([]int) returns a pointer to a newly allocated, zeroed slice structure, that is, a pointer to a nil slice value.
	i1 := new([]int) // 返回一个指针指向空的slice(nil slice)
	i2 := make([]int, 10)

	fmt.Println(i1, i2)

	// Remember that make applies only to maps, slices and channels and does not return a pointer.
	//  To obtain an explicit pointer allocate with new or take the address of a variable explicitly.

	// make 只用于 map slice channel 不返回指针，初始化内容 返回引用本身
	// new 返回指针 不初始化内容
}
