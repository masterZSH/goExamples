package main

import (
	"context"
	"fmt"
	"time"
)

type myKey string

func main() {
	// ctx := context.Background()
	// method1(ctx)
	method3()
}

func method1(ctx context.Context) {
	ctx = context.WithValue(ctx, myKey("testKey"), "Go")
	method2(ctx)
}

func method2(ctx context.Context) {
	fmt.Print(ctx.Value(myKey("testKey")))
	// 2s后超时
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)

	defer cancel()

	select {
	case <-time.After(10 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}
}

func method3() {
	// 内部一个协程 持续向通道写入数字 1 2 3 4 ...
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {

				// 如果接受到结束通道的数据 结束
				case <-ctx.Done():
					return
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}
