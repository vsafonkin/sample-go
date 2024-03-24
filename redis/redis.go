package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var client *redis.Client

func InitClient() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func SetValue(ctx context.Context, key string, val interface{}) error {
	return client.Set(ctx, key, val, 0).Err()
}

func GetValue(ctx context.Context, key string) (interface{}, error) {
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return val, nil
}
