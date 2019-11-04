package main

import (
	"bytes"
	"fmt"
)

func main() {
	// 读写长度未知的bytes最好使用Buffer结构体
	var r *bytes.Buffer = new(bytes.Buffer)
	fmt.Printf("%p", r)
	var buffer bytes.Buffer
	for {
		if s, ok := getNextString(); ok { //method getNextString() not shown here
			buffer.WriteString(s)
		} else {
			break
		}
	}
	fmt.Print(buffer.String(), "\n")

	// 把一个缓存 buf 分片成两个 切片：
	// 第一个是前 n 个 bytes，后一个是剩余的，用一行代码实现。
	var buf = []byte{'a', 'b', 'c'}
	var bufArr = []([]byte){buf[:1], buf[1:]}
	fmt.Print(bufArr)

	x := make([]int, 3)
	fmt.Printf("%p\n", &x) //不变 0xc000004480
	//y := make([]int, 3)
	y := &x
	fmt.Print(*y) //不变 0xc0000044c0
	fmt.Printf("%p\n", y)
	x[0] = 1
	fmt.Println(x)  //[1 0 0]
	fmt.Println(*y) //[1 0 0]

}

func getNextString() (string, bool) {
	return "234", false
}
