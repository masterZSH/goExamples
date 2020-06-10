package main

import "fmt"

func main() {
	f()
	// f()
	// g()
	// h()
	// call defer h()
	// call defer g()
	// call defer f()
	// 最先defer的最后执行
}

func f() {
	defer func() {
		fmt.Println("call defer f()")
	}()
	fmt.Println("f()")
	g()
}

func g() {
	defer func() {
		fmt.Println("call defer g()")
	}()
	fmt.Println("g()")
	h()
}

func h() {
	defer func() {
		fmt.Println("call defer h()")
	}()
	fmt.Println("h()")
}
