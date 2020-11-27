package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	ah "github.com/masterZSH/rservices/auth"
	"github.com/masterZSH/rservices/user"

	"github.com/docker/libkv/store"

	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/share"
)

var (
	addr     = flag.String("addr", "localhost:8972", "server address")
	etcdAddr = flag.String("etcd", "127.0.0.1:2379", "etcd address")
	basePath = flag.String("basePath", "zsh", "base path")
)

func getDefaultOptions() client.Option {
	option := client.DefaultOption
	option.Heartbeat = true
	option.HeartbeatInterval = time.Second
	// 设置超时时间
	option.ConnectTimeout = 10 * time.Second
	return option
}

func authConfig() *store.Config {
	// etcd3 auth
	return &store.Config{
		Username: "zsh",
		Password: "123456",
	}
}

func main() {
	flag.Parse()
	getUser()
}

func getUser() {
	option := getDefaultOptions()
	cg := authConfig()
	d := client.NewEtcdV3Discovery(*basePath, "User", []string{*etcdAddr}, cg)
	// 路由模式-随机选择
	xclient := client.NewXClient("User", client.Failtry, client.RandomSelect, d, option)
	xclient.Auth("")
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
		// token
		if err != nil && err.Error() == "invalid token" {
			fmt.Println("---------")
			// 获取token
			tk, err := getToken()
			if err != nil {
				continue
			}
			xclient.Auth(tk)
			err = xclient.Call(ctx, "GetUser", args, reply)
		}
		if err != nil {
			// debug.PrintStack()
			fmt.Printf("failed to call: %v", err)
			continue
		}
		log.Printf("get id : %d user %+v\n", args.UserID, reply.U)
		time.Sleep(1e9)
	}
}

func getToken() (token string, err error) {
	option := client.DefaultOption
	option.Heartbeat = true
	option.HeartbeatInterval = time.Second
	option.ConnectTimeout = 10 * time.Second
	cg := authConfig()
	d := client.NewEtcdV3Discovery(*basePath, "Auth", []string{*etcdAddr}, cg)
	// 路由模式-随机选择
	xclient := client.NewXClient("Auth", client.Failtry, client.RandomSelect, d, option)
	defer xclient.Close()
	args := &ah.AuthArgs{}
	reply := &ah.AuthReply{}
	// 同步 异步使用Go
	err = xclient.Call(context.Background(), "GetToken", args, reply)
	if err != nil {
		fmt.Printf("fail to get token :%+v", err)
		return "", err
	}
	return reply.Token, nil
}
