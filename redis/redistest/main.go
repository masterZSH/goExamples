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
		Addr: mr.Addr(),
	})
	ctx := context.Background()
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
