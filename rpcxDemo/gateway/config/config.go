package config

import (
	"time"

	"github.com/docker/libkv/store"
	"github.com/smallnest/rpcx/client"
)

type BaseConfig struct {
	Addr     string
	EtcdAddr string
	BasePath string
}

func AuthConfig() *store.Config {
	// etcd3 auth
	return &store.Config{
		Username: "zsh",
		Password: "123456",
	}
}

func DefaultConfig() BaseConfig {
	return BaseConfig{
		Addr:     "localhost:8972",
		EtcdAddr: "127.0.0.1:2379",
		BasePath: "zsh",
	}
}

func GetDefaultOptions() client.Option {
	option := client.DefaultOption
	option.Heartbeat = true
	option.HeartbeatInterval = time.Second
	// 设置超时时间
	option.ConnectTimeout = 10 * time.Second
	return option
}
