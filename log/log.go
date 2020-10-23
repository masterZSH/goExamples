package main

import (
	"log/log"
)

func main() {
	// out, _ := os.OpenFile("1.log", os.O_CREATE|os.O_APPEND, 0666)
	l := log.NewLogger(&log.Out{
		BaseDir: "D:\\log\\",
	})
	// fields 字段
	fields := make(map[string]interface{})
	fields["name"] = "z"
	l.WithFields(fields).Info("测试")

}
