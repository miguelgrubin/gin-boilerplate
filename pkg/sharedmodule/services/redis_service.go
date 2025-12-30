package services

import (
	"context"
	"time"

	redis "github.com/redis/go-redis/v9"
)

type RedisServiceInterface interface {
	Set(key string, value any, expiration time.Duration) error
	Get(key string) (string, error)
	Has(key string) (bool, error)
	Del(key string) error
}

type RedisService struct {
	config RedisConfig
	rdb    *redis.Client
	ctx    context.Context
}

func NewRedisService(c RedisConfig) *RedisService {
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Address,
		Password: c.Password,
		DB:       c.DB,
	})
	ctx := context.Background()

	return &RedisService{
		config: c,
		rdb:    rdb,
		ctx:    ctx,
	}
}

func (r *RedisService) Set(key string, value interface{}, expiration time.Duration) error {
	return r.rdb.Set(r.ctx, key, value, expiration).Err()
}

func (r *RedisService) Get(key string) (string, error) {
	return r.rdb.Get(r.ctx, key).Result()
}

func (r *RedisService) Has(key string) (bool, error) {
	result, err := r.rdb.Exists(r.ctx, key).Result()
	if err != nil {
		return false, err
	}
	return result > 0, nil
}

func (r *RedisService) Del(key string) error {
	return r.rdb.Del(r.ctx, key).Err()
}

var _ RedisServiceInterface = &RedisService{}
