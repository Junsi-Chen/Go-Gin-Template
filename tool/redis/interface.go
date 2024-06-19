package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

func NewIRedis() IRedis {
	return &Redis{
		rdb: GetRdb(),
	}
}

type IRedis interface {
	Get(key string) string
	Set(key, value string, expiration time.Duration) error
	BFAdd(key, value string) error
	BFExists(key, value string) bool
}
type Redis struct {
	rdb *redis.Client
}

func (r *Redis) Get(key string) string {
	value, _ := r.rdb.Get(context.Background(), key).Result()
	return value
}

func (r *Redis) Set(key, value string, expiration time.Duration) error {
	return r.rdb.Set(context.Background(), key, value, expiration).Err()
}

func (r *Redis) BFAdd(key, value string) error {
	return r.rdb.BFAdd(context.Background(), key, value).Err()
}

func (r *Redis) BFExists(key, value string) bool {
	exist, _ := r.rdb.BFExists(context.Background(), key, value).Result()
	return exist
}
