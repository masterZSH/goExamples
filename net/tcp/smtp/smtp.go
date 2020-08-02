package main

import (
	"log"
	"net/smtp"
)

// 简易发送邮件demo
func main() {
	// 授权密码获取： qq邮箱-开启pop3/smtp服务时获取
	auth := smtp.PlainAuth("castinue@qq.com", "castinue@qq.com", "授权密码", "smtp.qq.com")
	err := smtp.SendMail(
		"smtp.qq.com:587",
		auth,
		"castinue@qq.com",
		[]string{"1044438932@qq.com"},
		[]byte("<div>123</div>"),
	)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
