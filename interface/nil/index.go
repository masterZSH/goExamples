package main

import (
	"fmt"
	"sync"
	"time"
)

type Instance struct {
	sync.Mutex
	v int
}

var wg sync.WaitGroup

func main() {
	i := new(Instance)

	for v := 0; v < 10; v++ {
		wg.Add(1)
		go safeSadd(i)
	}
	wg.Wait()
}

func safeSadd(i *Instance) {
	i.Lock()
	i.v++
	fmt.Printf("v is %d\n", i.v)
	i.Unlock()
	time.Sleep(time.Second)
	wg.Done()
}
