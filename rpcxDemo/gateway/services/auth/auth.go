package auth

import (
	"context"
	"fmt"
	"main/config"
	"time"

	ah "github.com/masterZSH/rservices/auth"

	"github.com/smallnest/rpcx/client"
)

func GetToken() (token string, err error) {
	option := client.DefaultOption
	option.Heartbeat = true
	option.HeartbeatInterval = time.Second
	option.ConnectTimeout = 10 * time.Second
	cg := config.AuthConfig()
	cfg := config.DefaultConfig()
	d := client.NewEtcdV3Discovery(cfg.BasePath, "Auth", []string{cfg.EtcdAddr}, cg)
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
