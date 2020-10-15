package main

import (
	"fmt"
	"os"
)

func main() {

	// userdir
	// 返回用户配置目录
	configDir, _ := os.UserConfigDir()
	fmt.Println(configDir)

	// 用户主页目录
	homeDir, _ := os.UserHomeDir()
	fmt.Println(homeDir)

	// 用户缓存目录
	cacheDir, _ := os.UserCacheDir()
	fmt.Println(cacheDir)
}
