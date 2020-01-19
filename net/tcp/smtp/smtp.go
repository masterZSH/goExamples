package main

import (
	"log"
	"net/smtp"
)

func main() {
	// qq发送的授权密码
	auth := smtp.PlainAuth("castinue@qq.com", "castinue@qq.com", "授权密码", "smtp.qq.com")
	// ynnmqzdklbrhdiei
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
