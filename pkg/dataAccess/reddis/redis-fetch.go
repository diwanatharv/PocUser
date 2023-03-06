package redisdb

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"

	"npm/pkg/config"
)

type Redis struct {
}

func (cache *Redis) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return config.RedisDatabase.Set(context.Background(), key, value, expiration)
}
func (cache *Redis) Get(key string) *redis.StringCmd {
	return config.RedisDatabase.Get(context.Background(), key)
}
