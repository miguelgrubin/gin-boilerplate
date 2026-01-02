package services

import (
	"context"
	"time"

	redis "github.com/redis/go-redis/v9"
)

type RedisService interface {
	Set(key string, value any, expiration time.Duration) error
	Get(key string) (string, error)
	Has(key string) (bool, error)
	Del(key string) error
}

type RedisServiceImp struct {
	config RedisConfig
	rdb    *redis.Client
	ctx    context.Context
}

func NewRedisService(c RedisConfig) *RedisServiceImp {
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Address,
		Password: c.Password,
		DB:       c.DB,
	})
	ctx := context.Background()

	return &RedisServiceImp{
		config: c,
		rdb:    rdb,
		ctx:    ctx,
	}
}

func (r *RedisServiceImp) Set(key string, value interface{}, expiration time.Duration) error {
	return r.rdb.Set(r.ctx, key, value, expiration).Err()
}

func (r *RedisServiceImp) Get(key string) (string, error) {
	return r.rdb.Get(r.ctx, key).Result()
}

func (r *RedisServiceImp) Has(key string) (bool, error) {
	result, err := r.rdb.Exists(r.ctx, key).Result()
	if err != nil {
		return false, err
	}
	return result > 0, nil
}

func (r *RedisServiceImp) Del(key string) error {
	return r.rdb.Del(r.ctx, key).Err()
}

var _ RedisService = &RedisServiceImp{}
