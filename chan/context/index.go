package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex
var goroutineID int = 1

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go f(ctx)
	go f1(ctx)

	time.Sleep(10 * time.Second)
	cancel()
	time.Sleep(3 * time.Second)
}

func f(ctx context.Context) {
	mu.Lock()
	ID := goroutineID
	goroutineID++
	mu.Unlock()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%d停止了...\n", ID)
			return
		default:
			fmt.Printf("%d进行中...\n", ID)
			time.Sleep(2 * time.Second)
		}
	}
}

func f1(ctx context.Context) {
	go f(ctx)
	go f(ctx)
}
