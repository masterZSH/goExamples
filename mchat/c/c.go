package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/masterZSH/mino"
	"github.com/schollz/peerdiscovery"
)

type Message struct {
	Type int
	Data []byte
}

var k mino.Key

func main() {
	discoveries, _ := peerdiscovery.Discover(peerdiscovery.Settings{
		Limit:     1,
		AllowSelf: true,
	})
	if len(discoveries) == 0 {
		return
	}
	d := discoveries[0]
	var port string
	if strings.Contains(string(d.Payload), "chat:") {
		port = strings.Split(string(d.Payload), "chat:")[1]
	}
	conn, err := net.Dial("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	go ReadMessage(conn)

	for {
	}

}

func ReadMessage(conn net.Conn) {
	data, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		return
	}
	fmt.Println("Message Received: ", string(data))
	var m Message
	json.Unmarshal(data, &m)
	fmt.Printf("%s", m.Data)
	arr := strings.Split(string(m.Data), "-")
	k, _ = mino.NewKey([]byte(arr[0]), []byte(arr[1]))

	content := []byte("666")

	da, _ := k.Encrypt(content)

	m1 := Message{
		Type: 2,
		Data: da,
	}

	bs1, _ := json.Marshal(m1)
	bs1 = append(bs1, '\n')
	conn.Write(bs1)
}
