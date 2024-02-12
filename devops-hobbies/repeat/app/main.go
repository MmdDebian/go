package main

import (
	"context"

	"github.com/redis/go-redis/v9"
)


func main(){
	client := redis.NewClient(&redis.Options{
		Addr: "",
		Password: "",
		DB: 0,
	})

	ctx := context.Background()

	client.Set(ctx , "name" , "ali",0)
}