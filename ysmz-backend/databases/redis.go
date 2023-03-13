package databases

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisInterface interface {
	Get(key string) (string, error)
	Set(key string, value string, expiration time.Duration) error
}

type RedisStruct struct {
	client *redis.Client
	ctx    context.Context
}

var redisDB *RedisStruct // singleton

func RedisDB() *RedisStruct {
	if redisDB == nil {
		return &RedisStruct{
			client: redis.NewClient(&redis.Options{
				Addr:     "localhost:6379",
				Password: "", // no password set
				DB:       0,  // use default DB
			}),
			ctx: context.Background(),
		}
	}
	return redisDB
}

func (r *RedisStruct) Get(key string) (string, error) {
	return r.client.Get(r.ctx, key).Result()
}

func (r *RedisStruct) Set(key string, value string, expiration time.Duration) error {
	return r.client.Set(r.ctx, key, value, expiration).Err()
}
