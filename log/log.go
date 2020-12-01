package main

import (
	"log/log"
)

func main() {
	// out, _ := os.OpenFile("1.log", os.O_CREATE|os.O_APPEND, 0666)
	l := log.New("D:\\log\\")
	// fields 字段
	l.Info("测试")
}
