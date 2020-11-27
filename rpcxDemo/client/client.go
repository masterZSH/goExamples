package main

import (
	"context"
	"flag"
	"log"
	"main/rpc/services/user"
	"time"

	"github.com/docker/libkv/store"

	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/share"
)

var (
	addr     = flag.String("addr", "localhost:8972", "server address")
	etcdAddr = flag.String("etcd", "127.0.0.1:2379", "etcd address")
	basePath = flag.String("basePath", "zsh", "base path")
)

func main() {
	flag.Parse()
	// d := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")

	// 多台服务器提供相同的服务
	// addr1 := "localhost:8972"
	// addr2 := "localhost:8973"

	// zookeeper  etcd
	// 设置自动心跳1s
	option := client.DefaultOption
	option.Heartbeat = true
	option.HeartbeatInterval = time.Second

	// 设置超时时间
	option.ConnectTimeout = 10 * time.Second

	// d := client.NewMultipleServersDiscovery([]*client.KVPair{{Key: addr1}, {Key: addr2}})

	// etcdv3 服务发现

	// Auth
	cg := &store.Config{
		Username: "zsh",
		Password: "123456",
	}

	d := client.NewEtcdV3Discovery(*basePath, "User", []string{*etcdAddr}, cg)

	// 路由模式-随机选择
	xclient := client.NewXClient("User", client.Failtry, client.RandomSelect, d, option)
	xclient.Auth("bearer abcdefg1234567")

	defer xclient.Close()

	args := &user.Args{
		UserID: 1,
	}

	for {
		reply := &user.Reply{}

		// 同步 异步使用Go
		ctx := context.WithValue(context.Background(), share.ReqMetaDataKey, make(map[string]string))

		err := xclient.Call(ctx, "GetUser", args, reply)

		// 广播模式 调用节点
		// err := xclient.Broadcast(context.Background(), "GetUser", args, reply)

		if err != nil {
			log.Fatalf("failed to call: %v", err)
		}

		log.Printf("get id : %d user %+v\n", args.UserID, reply.U)
		time.Sleep(1e9)
	}

}
