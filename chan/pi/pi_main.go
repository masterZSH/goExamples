package main

import (
	"flag"
	"fmt"
	"math"
	"runtime"
	"time"
)

var (
	h, d    bool
	c, n, l int
)

func init() {
	// 帮助信息
	flag.BoolVar(&h, "h", false, "帮助信息")
	// 使用cpu核心数量 默认1
	flag.IntVar(&c, "c", 1, "使用cpu核心数量")
	// 需要计算的次数 默认1
	flag.IntVar(&n, "n", 1, "需要计算的次数,越大越精确同时消耗的时间越长")
	// 打印消耗时间
	flag.BoolVar(&d, "d", false, "打印消耗时间")
	// 缓冲长度 默认10
	flag.IntVar(&l, "l", 10, "缓冲长度")
}

// CalcPi 计算Pi圆周率主方法
// times int 计算次数
func CalcPi(times int) float64 {
	// 创建一个channel存储计算值 缓冲长度10
	ch := make(chan float64, l)
	// cpu多核心运算
	for i := 0; i < c; i++ {
		go term(ch, i*times/c, (i+1)*times/c)
	}
	// 累加channel存储的值
	result := 0.0
	for i := 0; i < c; i++ {
		result += <-ch
	}
	return result
}

// term 计算主方法
func term(ch chan float64, start, end int) {
	result := 0.0
	for i := start; i < end; i++ {
		x := float64(i)
		result += 4 * (math.Pow(-1, x) / (2.0*x + 1.0))
	}
	ch <- result
}

func main() {
	flag.Parse()
	if h {
		// 打印信息
		flag.PrintDefaults()
	}
	// 开始时间
	start := time.Now()
	// 设置cpu使用数量
	runtime.GOMAXPROCS(c)
	// 输出计算值
	fmt.Println(CalcPi(n))
	// 结束时间
	end := time.Now()
	// 消耗时间
	delta := end.Sub(start)
	// 打印计算耗时
	if d {
		fmt.Printf("计算总耗时: %s\n", delta)
	}
}
