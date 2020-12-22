package main

import (
	"context"
	"fmt"
	"time"
)

type myKey string

func main() {
	ctx := context.Background()
	method1(ctx)
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
