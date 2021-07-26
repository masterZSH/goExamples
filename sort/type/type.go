package main

import (
	"fmt"
	"sync"
)

type keyValues struct {
	key    string
	values []string
}

type headerSorter struct {
	kvs []keyValues
}

var headerSorterPool = sync.Pool{
	New: func() interface{} { return new(headerSorter) },
}

func main() {
	// var st interface{} = new(keyValues)
	// fmt.Println(reflect.TypeOf(st).String())

	// hs := headerSorterPool.Get().(*headerSorter)
	// s,_:= st.(*keyValues)
	// fmt.Println(reflect.TypeOf(s).String())
	// s1:=(interface{})(s)
	// fmt.Println(reflect.TypeOf(s1).String())
	// fmt.Printf("%v",st);
	// fmt.Printf("%v",s);
	// fmt.Print(headerSorterPool.Get().(*headerSorter))
	t = 1
	returnInt()
	t = "123"
	returnString()

}

var t interface{}

func returnInt() int {
	fmt.Print(t)
	return t.(int)
}

func returnString() string {
	fmt.Print(t)
	return t.(string)
}
