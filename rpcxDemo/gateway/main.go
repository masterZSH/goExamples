package main

import (
	"main/routers"
)

func main() {
	// gin 工程实例
	router := routers.Router
	// 路由初始化
	routers.RegisterRouter()
	router.Run(":8656")

}
