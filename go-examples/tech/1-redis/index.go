package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()

	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()

	if err != nil {
		panic(err)
	}

	fmt.Println("key", val)
}
