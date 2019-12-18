package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

const NCPU = 4

// pi计算 随着计算越来越精确
func pi(ch chan float64, end int) {
	var r float64
	for k := 0; k < end; k++ {
		i := float64(k)
		item := 4 * (math.Pow(-1, i) / (2.0*i + 1.0))
		r += float64(item)
	}
	ch <- r
}

func CalculatePi(end int) float64 {
	ch := make(chan float64)
	for i := 0; i < NCPU; i++ {
		go term(ch, i*end/NCPU, (i+1)*end/NCPU)
	}
	result := 0.0
	for i := 0; i < NCPU; i++ {
		result += <-ch
	}
	return result
}

func term(ch chan float64, start, end int) {
	result := 0.0
	for i := start; i < end; i++ {
		x := float64(i)
		result += 4 * (math.Pow(-1, x) / (2.0*x + 1.0))
	}
	ch <- result
}

func main() {
	start := time.Now()
	runtime.GOMAXPROCS(NCPU)
	fmt.Println(CalculatePi(50000000))
	end := time.Now()
	delta := end.Sub(start)
	// 对比时间 pi_calc 500ms左右 这个190ms提升很明显
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}
