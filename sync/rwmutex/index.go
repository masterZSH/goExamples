package main

import (
	"fmt"
	"sync"
	"time"
)

type A struct {
	sync.RWMutex
	v int
}

var wg sync.WaitGroup

func main() {
	a := new(A)
	wg.Add(4)
	go writeValue(a)
	go readValue(a)
	go readValue(a)
	go readValue(a)

	wg.Wait()
}

func readValue(a *A) {
	a.RLock()
	fmt.Printf("get v:%d\n", a.v)
	a.RUnlock()
	wg.Done()
}

func writeValue(a *A) {
	a.Lock()
	fmt.Printf("set v:%d\n", a.v)
	a.v = 1
	time.Sleep(time.Second)
	a.Unlock()
	wg.Done()
}

// output
// get v:0
// get v:0
// set v:0
// get v:1
