package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

// 聊天服务版本号
const version string = "1.0.0"

// Client 客户端
type Client struct {
	conn net.Conn
	name string
	msg  chan *Message
}

// NewClient 创建客户端
func NewClient(conn net.Conn) *Client {
	return &Client{
		conn: conn,
		name: "",
		msg:  make(chan *Message),
	}
}

// Message 消息结构
type Message struct {
	time    time.Time
	content string
	client  *Client
}

// NewMessage 新建消息
func NewMessage(client *Client, content string) *Message {
	return &Message{
		client:  client,
		time:    time.Now(),
		content: content,
	}
}

var (
	// 端口号
	p int
	// 是否展示帮助信息
	h bool
	// 版本号
	v bool
	// 输出流
	out = os.Stdout
	// 进入聊天室通道
	enter = make(chan *Client)
	// 离开聊天室通道
	leave = make(chan *Client)
	// 广播消息通道
	messages = make(chan *Message)
)

func init() {
	flag.IntVar(&p, "p", 6666, "监听端口号")
	flag.BoolVar(&h, "h", false, "帮助信息")
	flag.BoolVar(&v, "v", false, "版本")
}

func main() {
	flag.Parse()
	if h {
		flag.PrintDefaults()
	}
	if v {
		fmt.Fprintf(out, "聊天服务版本号:%s\n", version)
		os.Exit(0)
	}
	run()
}

// 服务主入口
func run() {
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", p))
	if err != nil {
		fmt.Fprintf(out, "%v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(out, "服务已启动，监听端口号[%d]\n", p)
	// 广播协程
	go broadCast()
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleConn(conn)
	}
}

// 广播
func broadCast() {
	clients := make(map[*Client]interface{})
	for {
		select {
		case message := <-messages:
			for client := range clients {
				if client.name != "" {
					client.msg <- message
				}
			}
		case client := <-enter:
			clients[client] = struct{}{}
		case client := <-leave:
			delete(clients, client)
			close(client.msg)
		}
	}
}

// 客户端连接处理
func handleConn(conn net.Conn) {
	// 客户端初始化
	var client *Client
	client = NewClient(conn)
	// 写入客户端协程
	go writeToClient(client)
	client.msg <- NewMessage(client, fmt.Sprintf("请输入名称："))
	// 客户端输入处理
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		content := scanner.Text()
		if client.name == "" {
			enterChatRoom(client, content)
			continue
		}
		messages <- NewMessage(client, fmt.Sprintf("[%s]:%s", client.name, content))
	}
	leaveChatRoom(client)
}

// 写入消息到客户端
func writeToClient(client *Client) {
	for msg := range client.msg {
		var output string
		if client.name == "" {
			output = fmt.Sprintf("%s\n", msg.content)
		} else {
			output = fmt.Sprintf("%s        %s\n", msg.content, msg.time.Format("2006-01-02 15:04:05"))
		}
		fmt.Fprintf(client.conn, output)
	}
}

// 进入聊天室
func enterChatRoom(client *Client, name string) {
	client.name = name
	enter <- client
	messages <- NewMessage(client, fmt.Sprintf("%s进入聊天室", client.name))
}

// 离开聊天室
func leaveChatRoom(client *Client) {
	leave <- client
	client.conn.Close()
	messages <- NewMessage(client, fmt.Sprintf("%s离开聊天室", client.name))
}
