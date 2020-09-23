package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {

	// 将字符串中 #{name} 特殊标识替换为params[name]中的值
	str := "#{name}你好，您可以点击链接#{ticket_url}进行索票，谢"
	re := regexp.MustCompile(`#{(\w+)}`)
	arr := re.FindAllStringSubmatch(str, -1)
	params := make(map[string]string)
	params["name"] = "crisnina"
	params["ticket_url"] = "https://ticket.com"
	content := str
	for _, item := range arr {
		fmt.Println(item)
		if param, ok := params[item[1]]; ok {
			content = strings.Replace(content, item[0], param, -1)
		}
	}
	fmt.Println(arr)
	fmt.Println(content)

}
