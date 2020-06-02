package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	now:=time.Now()
	urls:=os.Args[1:]
	ch := make(chan string,10)
	// 并发请求
	for _,url := range urls{
		go fetch(url,ch)
	}
	
	for i:=0; i< len(urls);i++ {
		fmt.Printf("%s\n",<-ch)
	}

	defer func ()  {
		close(ch)
	}()
	// receive(ch)
	// 计算总花费时间
	fmt.Printf("总花费时间%.2fs",time.Since(now).Seconds())
}


func fetch(url string,ch chan string)  {
	// 缺省参数
	if !strings.HasPrefix(url,"http://"){
		url = "http://" + url
	}
	now:=time.Now()
	resp, err := http.Get(url)
	if err != nil{
		ch<-fmt.Sprintf("错误%v",err)
		return 
	}
	
	defer func ()  {
		resp.Body.Close()
	}()

	nums,err := io.Copy(ioutil.Discard,resp.Body)
	if err != nil{
		ch<-fmt.Sprintf("错误%v",err)
		return 
	}
	speedTime:=time.Since(now)
	str:= fmt.Sprintf("读取[%d]个字节，花费%.2fs时间",nums,speedTime.Seconds())
	ch<-str	
}


func receive(ch chan string){
	for s:= range ch{
		fmt.Printf("%s",s)
	}
}