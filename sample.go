package main

import (
	"context"
	"fmt"
	"github.com/vsafonkin/sample-go/redis"
)

func main() {
	redis.InitClient()

	key := "bob"
	ctx := context.Background()
	err := redis.SetValue(ctx, key, 789)
	if err != nil {
		fmt.Println("set error:", err)
	}

	val, err := redis.GetValue(ctx, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
}
