package main

import (
	"bytes"
	"fmt"
)

func main() {
	// closeSelect()

	// RunTimePanic()
	// RunTimePanic2()
	// RunTimePanic3()
	// RunTimePanic4()

	ChannelIsClosed()
}

func closeSelect() {
	done := make(chan struct{})
	data := bytes.NewBuffer([]byte("123"))
	var body []byte
	go func() {
		var buf bytes.Buffer
		buf.ReadFrom(data)
		body = buf.Bytes()
		close(done)
	}()
	select {
	case <-done:
		fmt.Printf("close down %s", body)
	}
}

func RunTimePanic() {
	receiveOnlyCh := make(chan<- struct{})
	receiveOnlyCh <- struct{}{}
	close(receiveOnlyCh)
}

func RunTimePanic2() {
	ch := make(chan struct{}, 1)
	ch <- struct{}{}
	close(ch)
	ch <- struct{}{}

}

// close nil channel will run-time panic
func RunTimePanic3() {
	var ch chan int
	close(ch)
}

func RunTimePanic4() {
	ch := make(chan int, 10)
	ch <- 1
	close(ch)
	close(ch)

}

func ChannelIsClosed() {
	ch := make(chan int, 10)
	ch <- 1
	close(ch)

	data, c := <-ch
	// c => is closed
	fmt.Println(data, c)
}
