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

func GetValues(ctx context.Context, pattern string) ([]string, error) {
	r, err := client.Do(ctx, "KEYS", pattern).Result()
	if err != nil {
		return nil, err
	}
	var out []string
	ri := r.([]interface{})
	for i := range ri {
		out = append(out, ri[i].(string))
	}

	return out, nil
}
