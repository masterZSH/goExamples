package main

import (
	"fmt"
	"sync"
	"time"
)

// RWMutex 读写锁

type A struct {
	sync.RWMutex
	v int
}

var wg sync.WaitGroup

var (
	rwLock sync.RWMutex
	num    int
)

func main() {
	a := new(A)
	wg.Add(6)
	// 修改v = 1
	go writeValue(a, 1)

	// 第二个写锁要等 其他读/写锁释放后可用
	// Lock locks rw for writing. If the lock is already locked for reading or writing, Lock blocks until the lock is available.
	go writeValue(a, 2)

	// readValue读取都是一个值1
	go readValue(a)
	go readValue(a)
	go readValue(a)
	go readValueAfterSecond(a) // 0
	wg.Wait()

}

func readValue(a *A) {
	a.RLock()
	fmt.Printf("get v:%d\n", a.v)
	a.RUnlock()
	wg.Done()
}

func readValueAfterSecond(a *A) {
	a.RLock()
	time.Sleep(time.Second)
	fmt.Printf("after a second get v:%d\n", a.v)
	a.RUnlock()
	wg.Done()
}

func writeValue(a *A, val int) {
	a.Lock()
	a.v = val
	fmt.Printf("set v:%d\n", a.v)
	time.Sleep(time.Second)
	a.Unlock()
	wg.Done()
}

// output
// get v:0
// get v:0
// set v:0
// get v:1
