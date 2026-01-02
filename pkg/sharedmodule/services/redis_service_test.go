package services_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/miguelgrubin/gin-boilerplate/pkg/sharedmodule/services"
	"github.com/stretchr/testify/suite"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/redis"
)

type RedisServiceTestSuite struct {
	suite.Suite
	ctx            context.Context
	redisContainer *redis.RedisContainer
}

func (suite *RedisServiceTestSuite) SetupTestSuite() {
	suite.ctx = context.Background()

	rc, _ := redis.Run(suite.ctx, "redis:8")
	suite.redisContainer = rc
}

func (suite *RedisServiceTestSuite) TearDownTestSuite() {
	testcontainers.CleanupContainer(suite.T(), suite.redisContainer)
}

func (ts *RedisServiceTestSuite) TestSet() {
	host, _ := ts.redisContainer.Host(ts.ctx)
	port, _ := ts.redisContainer.MappedPort(ts.ctx, "6379")

	redisService := services.NewRedisService(services.RedisConfig{
		Address:  fmt.Sprintf("%s:%s", host, port.Port()),
		Password: "",
		DB:       0,
	})
	err := redisService.Set("test-key", "test-value", time.Minute.Abs())

	ts.NoError(err)
}

func (ts *RedisServiceTestSuite) TestGet() {
	host, _ := ts.redisContainer.Host(ts.ctx)
	port, _ := ts.redisContainer.MappedPort(ts.ctx, "6379")

	redisService := services.NewRedisService(services.RedisConfig{
		Address:  fmt.Sprintf("%s:%s", host, port.Port()),
		Password: "",
		DB:       0,
	})
	redisService.Set("test-key", "test-value", time.Minute.Abs())
	value, err := redisService.Get("test-key")

	ts.NoError(err)
	ts.Equal("test-value", value)
}

func (ts *RedisServiceTestSuite) TestDel() {
	host, _ := ts.redisContainer.Host(ts.ctx)
	port, _ := ts.redisContainer.MappedPort(ts.ctx, "6379")

	redisService := services.NewRedisService(services.RedisConfig{
		Address:  fmt.Sprintf("%s:%s", host, port.Port()),
		Password: "",
		DB:       0,
	})
	redisService.Set("test-key", "test-value", time.Minute.Abs())
	err := redisService.Del("test-key")

	ts.NoError(err)
}

func (ts *RedisServiceTestSuite) TestHas() {
	host, _ := ts.redisContainer.Host(ts.ctx)
	port, _ := ts.redisContainer.MappedPort(ts.ctx, "6379")
	redisService := services.NewRedisService(services.RedisConfig{
		Address:  fmt.Sprintf("%s:%s", host, port.Port()),
		Password: "",
		DB:       0,
	})

	notPresent, err := redisService.Has("test-key-not-exists")
	ts.NoError(err)
	redisService.Set("test-key", "test-value", time.Minute.Abs())
	isPresent, err := redisService.Has("test-key")

	ts.NoError(err)
	ts.True(isPresent)
	ts.False(notPresent)
}

func TestRedisService(t *testing.T) {
	dbConnSuite := new(RedisServiceTestSuite)
	dbConnSuite.SetupTestSuite()
	suite.Run(t, dbConnSuite)
	dbConnSuite.TearDownTestSuite()
}
