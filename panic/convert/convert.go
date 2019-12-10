package main

import (
	"fmt"
	"math"
)

// 写一个 ConvertInt64ToInt 函数把 int64 值转换为 int 值，
// 如果发生错误（提示：参见 4.5.2.1 节）就 panic。
// 然后在函数 IntFromInt64 中调用这个函数并 recover，
// 返回一个整数和一个错误。请测试这个函数！
func ConvertInt64ToInt(source int64) int {
	if math.MinInt32 <= source && source <= math.MaxInt32 {
		return int(source)
	}
	panic(fmt.Sprintf("%d is out of the int32 range", source))
}

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Print(e)
		}
	}()
	// fmt.Print(ConvertInt64ToInt(111))
	fmt.Print(ConvertInt64ToInt(1111111111111))
}
