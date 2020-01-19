package main

import "fmt"

import "time"

func main() {
	// 创建
	ch := make(chan int, 10)
	// 检测通道是否关闭
	// for {
	// 	if input, open := <-ch; !open {
	// 		break
	// 	}
	// 	fmt.Printf("%s", input)
	// }
	// 信号量模式
	go func() {
		// 如果希望程序一直阻塞，在匿名函数中省略 ch <- 1即可。
		ch <- 1
	}()
	otherThings()
	<-ch

	// 超时模板

	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(time.Second)
		timeout <- true
	}()
	select {
	case <-ch:
		fmt.Printf("%s", "get value from channel")
	case <-timeout:
		fmt.Printf("timeout")
	}

	// 输入通道和输出通道代替锁

}

type Task struct {
	name string
}

// Worker 使用输入通道和输出通道代替锁：
func Worker(in, out chan *Task) {
	for {
		t := <-in
		process(t)
		out <- t
	}
}

func process(t *Task) {
	fmt.Print(t.name)
}

func otherThings() {
	fmt.Print("sss")
}
