package main

import (
	"context"
	"flag"
	"log"
	"main/rpc/services/user"
	"time"

	"github.com/smallnest/rpcx/client"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()
	// d := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")

	// 多台服务器提供相同的服务
	addr1 := "localhost:8972"
	addr2 := "localhost:8973"

	// zookeeper  etcd
	// 设置自动心跳1s
	option := client.DefaultOption
	option.Heartbeat = true
	option.HeartbeatInterval = time.Second
	d := client.NewMultipleServersDiscovery([]*client.KVPair{{Key: addr1}, {Key: addr2}})
	xclient := client.NewXClient("User", client.Failtry, client.RandomSelect, d, option)
	defer xclient.Close()

	args := &user.Args{
		UserID: 1,
	}

	for {
		reply := &user.Reply{}

		// 同步 异步使用Go
		// err := xclient.Call(context.Background(), "GetUser", args, reply)

		// 广播模式
		err := xclient.Broadcast(context.Background(), "GetUser", args, reply)

		if err != nil {
			log.Fatalf("failed to call: %v", err)
		}

		log.Printf("get id : %d user %+v\n", args.UserID, reply.U)
		time.Sleep(1e9)
	}

}
