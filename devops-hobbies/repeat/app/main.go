package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)


func main(){
	client := redis.NewClient(&redis.Options{
		Addr: "",
		Password: "",
		DB: 0,
	})

        ctx := context.Background()

	// client.Set(ctx , "name" , "ali",0)

	session := map[string]string{"name":"mamad","username":"mohamad20"}

	for k,v := range session {
		err := client.HSet(ctx , "user-session:123",k,v).Err()

		if err != nil {
			panic(err)
		}
	}

	userSession := client.HGetAll(ctx,"user-session:123").Val()

	fmt.Println(userSession)

}
