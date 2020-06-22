package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

var (
	// Out 输出 默认为标准输出
	Out = os.Stdout
	// In 输入 默认为标准输入
	In = os.Stdin
	// Addr ip地址
	Addr string
	// Port 端口
	Port int
)

type Writer struct {
}

func Parse(str string) string {
	// strings.Replace(str, "\n", "", -1)
	arr := strings.Split(str, " ")
	// 结尾是\r\n
	var returnStr string
	returnStr += fmt.Sprintf("*%d\r\n", len(arr))
	for _, item := range arr {
		returnStr += fmt.Sprintf("$%d\r\n", len(item))
		returnStr += fmt.Sprintf("%s\r\n", item)
	}
	return returnStr
}

func init() {
	flag.StringVar(&Addr, "h", "", "host")
	flag.IntVar(&Port, "p", 6379, "port")
}

func main() {
	flag.Parse()
	conn, _ := net.Dial("tcp", fmt.Sprintf("%s:%d", Addr, Port))

	ch := make(chan string)
	end := make(chan bool)
	// fmt.Fprintf(conn, "*2\r\n$4\r\nINFO\r\n$3\r\nALL\r\n")

	go input(ch)
	go write(conn, ch)
	go readConn(conn, end)

	// 同步阻塞
	<-end
	// reader := bufio.NewReader(conn)
	// lin, _, _ := reader.ReadLine()
	defer conn.Close()
}

func input(ch chan string) {
	scanner := bufio.NewScanner(In)
	for scanner.Scan() {
		resp := Parse(scanner.Text())
		ch <- resp
	}
}

func write(conn net.Conn, ch chan string) {
	for resp := range ch {
		conn.Write([]byte(resp))
	}
}

//
func readConn(conn net.Conn, end chan bool) {
	connScanner := bufio.NewScanner(conn)
	for connScanner.Scan() {
		str := connScanner.Text()
		fmt.Println(str)
	}
	end <- true
}
