package main

import (
	"flag"
	"fmt"
	"io"
	"net"
)

func main() {
	flag.Parse()
	if flag.NArg() != 2 {
		panic("need host port")
	}

	data := make([]uint8, 4096)
	hostAndPort := fmt.Sprintf("%s:%s", flag.Arg(0), flag.Arg(1))
	// init server
	conn, err := net.Dial("tcp", hostAndPort) // tcp ipv4
	checkError(err, "conn error")
	io.WriteString(conn, "GET / \n")
	checkError(err, "conn error")
	read := true
	count := 0
	// 读取服务器的响应
	for read {
		count, err = conn.Read(data)
		read = (err == nil)
		fmt.Printf(string(data[0:count]))
	}
	conn.Close()
}

// 错误处理 panic
func checkError(error error, info string) {
	if error != nil {
		panic("ERROR: " + info + " " + error.Error()) // terminate program
	}
}
