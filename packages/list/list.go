package main

import (
	"container/list"
	"fmt"
	"unsafe"
)

func main() {
	intBytes()
}

// 通过使用 unsafe 包中的方法来测试你电脑上一个整型变量占用多少个字节。
func intBytes() {
	var i int
	fmt.Print(unsafe.Sizeof(i)) // 8个字节  64位的机器
}

// 使用 container/list 包实现一个双向链表，
// 将 101、102 和 103 放入其中并打印出来。
func doubleLinkedList() {
	// 双链表操作
	var l = list.New()
	l.PushBack(101)
	l.PushBack(102)
	l.PushBack(103)

	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value)
	}
	// l.PushFront(3)
	// for i := l.Front(); i != nil; i = i.Next() {
	// 	fmt.Print(i.Value)
	// }
}
