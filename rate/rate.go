package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"golang.org/x/time/rate"
)

type LimitReader struct {
	r       io.Reader
	limiter *rate.Limiter
}

func (r *LimitReader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	if err != nil {
		return
	}
	if r.limiter != nil {
		r.limiter.Wait(context.Background())
	}
	return n, err
}

func NewLimiterReader(r io.Reader, limit int64) io.Reader {

	// 每秒限制数
	if limit > 0 {
		return &LimitReader{
			r: r,

			// 每秒限制limit个event 令牌桶容量1
			limiter: rate.NewLimiter(rate.Limit(limit), 1),
		}
	}
	return r
}

func main() {

	// 大约50
	file, err := os.Open("t")
	if err != nil {
		panic(err)
	}

	// 1s读取10个
	rl := NewLimiterReader(file, 10)
	buf := make([]byte, 10)

	start := time.Now()
	for {
		n, err := rl.Read(buf)
		// n, err := file.Read(buf)
		if err != nil {
			break
		}
		fmt.Printf("%s\n", buf[:n])
	}

	end := time.Now()
	fmt.Println(end.Sub(start))

}
