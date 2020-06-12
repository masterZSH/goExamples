package main

import (
	"io"
	"log"
	"net"
	"os"
)

var (
	done   = make(chan struct{}) // 连接结束
	output = os.Stdout           // 输出，默认标准输出,可以修改为输出到文件
)

func main() {
	conn, err := net.Dial("tcp", "localhost:6666")
	if err != nil {
		log.Fatal(err)
	}
	go read(conn)
	write(conn, os.Stdin)
	conn.Close()

	// 解除阻塞
	<-done
}

// 从连接读取消息
func read(conn net.Conn) {
	io.Copy(output, conn)
	log.Println("服务器连接断开...")
	done <- struct{}{}
}

// 写入
func write(conn net.Conn, src io.Reader) {
	if _, err := io.Copy(conn, src); err != nil {
		log.Fatal(err)
	}
}
