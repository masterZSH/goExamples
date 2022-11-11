package main

import (
	"context"
	"fmt"
	"log"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis/v9"
)

func main() {
	mr, err := miniredis.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mr.Addr())
	client := redis.NewClient(&redis.Options{
		// Addr: "123.123.123.123:80",
		Addr: mr.Addr(),
	})
	ctx := context.Background()
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		// 连接失败
		log.Println(err)
	}
	fmt.Println(pong, err)
	err = client.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		log.Fatal(err)
	}

	val, err := client.Get(ctx, "key").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val)

}
