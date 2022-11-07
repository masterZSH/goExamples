package main

import (
	"context"
	"fmt"
)

func main() {
	ctx1 := context.Background()
	key := "foo"
	val := "bar"
	ctx2 := context.WithValue(ctx1, key, val)
	fmt.Println(ctx2.Value(key))
}
