package main

import (
	"context"
	"fmt"
	"time"
)

func ToChanTimedContext[T any](ctx context.Context, d time.Duration, message T, c chan<- T) (written bool) {
	ctx, cancel := context.WithTimeout(ctx, d)
	defer cancel()
	select {
	case c <- message:
		return true
	case <-ctx.Done():
		return false
	}
}

func ToChanTimedTimer[T any](d time.Duration, message T, c chan<- T) (written bool) {
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case c <- message:
		return true
	case <-t.C:
		return false
	}
}

func main() {
	ch := make(chan string, 1)
	d := 20 * time.Second
	msg := "msg"
	written := ToChanTimedContext[string](
		context.Background(),
		d,
		msg,
		ch,
	)
	fmt.Println(written) // true

	ch = make(chan string)
	d = 5 * time.Second
	msg = "msg"
	written = ToChanTimedContext[string](
		context.Background(),
		d,
		msg,
		ch,
	)
	fmt.Println(written) // false

	written = ToChanTimedTimer(
		d,
		msg,
		ch,
	)
	fmt.Println(written) // false

	ch = make(chan string, 1)

	written = ToChanTimedTimer(
		d,
		msg,
		ch,
	)
	fmt.Println(written) // true
}
