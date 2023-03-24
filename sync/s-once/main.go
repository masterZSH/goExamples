package once

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Once struct {
	mu   sync.Mutex // guards
	done uint32
}

func (e *Once) Do(f func()) {
	val := atomic.LoadUint32(&e.done)
	if val == 0 {
		e.mu.Lock()
		atomic.StoreUint32(&e.done, 1)
		f()
		e.mu.Unlock()
	}
}

// func main() {
// 	var once Once
// 	for i := 0; i < 10; i++ {
// 		go once.Do(testFunc)
// 	}
// 	time.Sleep(time.Second)

//}

func testFunc() {
	fmt.Println(1)
}
