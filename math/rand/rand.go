package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	l sync.Mutex
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func IntN(n int) int {
	l.Lock()
	defer l.Unlock()
	return r.Intn(n)
}

func Float64() float64 {
	l.Lock()
	defer l.Unlock()
	return r.Float64()
}

func Float32() float32 {
	l.Lock()
	defer l.Unlock()
	return r.Float32()
}

func Shuffle(n int, swap func(i, j int)) {
	l.Lock()
	defer l.Unlock()
	r.Shuffle(n, swap)
}

func main() {
	i := IntN(100)
	fmt.Println(i)

	t := []int{1, 2, 3, 4, 5}

	Shuffle(len(t), func(i, j int) {
		t[i], t[j] = t[j], t[i]
		fmt.Println(i, j)
	})
	fmt.Println(t)

}
