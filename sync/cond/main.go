package main

import (
	"fmt"
	"sync"
	"time"
)

// Cond
//
// 多个goroutines等待、1个goroutine通知事件发生。
func main() {

	var l sync.Mutex
	var f bool
	c := sync.NewCond(&l)

	var wg sync.WaitGroup
	wg.Add(3)

	// goroutine1
	go func() {

		c.L.Lock()

		for !f {
			fmt.Printf("goroutine1 wait \n")
			c.Wait()
		}
		fmt.Println("goroutine1 ", f)
		c.L.Unlock()
		wg.Done()

	}()

	// goroutine2
	go func() {
		c.L.Lock()

		for !f {
			fmt.Printf("goroutine2 wait \n")
			c.Wait()
		}
		fmt.Println("goroutine2 ", f)

		c.L.Unlock()
		wg.Done()

	}()

	// goroutine3
	go func() {
		c.L.Lock()

		for !f {
			fmt.Printf("goroutine3 wait \n")
			c.Wait()
		}
		fmt.Println("goroutine3 ", f)
		c.L.Unlock()

		wg.Done()
	}()

	// 模拟等待资源释放
	time.Sleep(2 * time.Second)

	// 资源已释放
	fmt.Println("ready")
	c.L.Lock()

	// 改变条件
	f = true

	// 广播所有的协程 检查条件是否满足即 !f == false 则执行后面的
	c.Broadcast()
	c.L.Unlock()

	wg.Wait()
}
