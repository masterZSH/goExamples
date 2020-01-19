package main

import "log"

func main() {
	protect(pa)
}

// 使用内建函数recover终止panic过程
func protect(g func()) {
	defer func() {
		log.Println("done")
		if x := recover(); x != nil {
			log.Printf("run time panic: %v", x)
		}
	}()
	log.Println("start")
	g()
}

func pa() {
	panic("error")
}
