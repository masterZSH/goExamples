package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Print(intN())
}

func intN() int {
	// 时间随机
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// returns [0,n)
	return r.Intn(2)
}
