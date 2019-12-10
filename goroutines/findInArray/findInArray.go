package main

import "fmt"

func main() {
	// int array
	var arr [100]int
	for i := 0; i < 100; i++ {
		arr[i] = i * i
	}
	// 开启10个协程查找数值
	for i:=0;i<10;i++{
		go findArr(256,arr[10*i:10*(i+1)])
	}

	// 死循环 避免main退出协程销毁
	for{
		
	}
}


func findArr(sr int, arr []int) int {
	found:=false
	for k, v := range arr {
		if v == sr{
			found = true
			fmt.Printf("find %d in key: %d\n",sr,k)
		}
		// fmt.Printf("k:%d,v:%d",k, v)
	}
	if !found{
		fmt.Println("not found")
	}
	return -1  // not found
}
