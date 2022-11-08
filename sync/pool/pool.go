package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

var (
	pool = sync.Pool{
		New: func() any {
			return make(map[string]int)
		},
	}

	p2 = sync.Pool{
		New: func() any {
			return new(bytes.Buffer)
		},
	}
)

func main() {
	// t()
	// defer func() {
	// 	t()
	// }()

	i := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
	}
	pool.Put(i)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		item1 := pool.Get().(map[string]int)
		fmt.Printf("%v\n", item1)
		wg.Done()
	}()
	go func() {
		item2 := pool.Get().(map[string]int)
		fmt.Printf("%v\n", item2)
		wg.Done()
	}()

	wg.Wait()

	closeCh := make(chan any)
	go func() {
		k := 0
		for k < 10 {
			Log(
				os.Stdout,
				"path",
				fmt.Sprintf("%s/%d\n", "/search?q=flowers", k),
			)
			k++
		}
		closeCh <- 1
	}()
	for {
		select {
		case <-closeCh:
			return
		}
	}

}

func Now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func Log(w io.Writer, key, val string) {
	b := p2.Get().(*bytes.Buffer)
	b.Reset()
	b.WriteString(Now())
	b.WriteByte(' ')
	b.WriteString(key)
	b.WriteByte('=')
	b.WriteString(val)
	w.Write(b.Bytes())
	p2.Put(b)
}
