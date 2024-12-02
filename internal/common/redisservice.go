package common

import (
	"fmt"
	"time"

	"github.com/cloudfresco/sc-dcsa/internal/config"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

/* from https://github.com/go-redis/redis/blob/master/command.go
type baseCmd struct {
	_args []interface{}
	err   error

	_readTimeout *time.Duration
}

type StringCmd struct {
	baseCmd

	val string
}

type StatusCmd struct {
	baseCmd

	val string
}

	Get(key string) *StringCmd
	Set(key string, value interface{}, expiration time.Duration) *StatusCmd

*/

// RedisIntf Interface to Redis commands
// All redis command to be called using this interface
type RedisIntf interface {
	Get(key string) (string, error)
	Set(key string, value interface{}, expiration time.Duration) error
}

// RedisService - Redis Pointer to redis
type RedisService struct {
	log         *zap.Logger
	RedisClient *redis.Client
}

// NewRedisService get connection to redis and create a RedisService struct
func NewRedisService(log *zap.Logger, redisOpt *config.RedisOptions) (*RedisService, error) {
	redisClient := redis.NewClient(&redis.Options{
		PoolSize:    10, // default
		IdleTimeout: 30 * time.Second,
		Addr:        redisOpt.Addr,
		Password:    "", // no password set
		DB:          0,  // use default DB
	})

	redisService := RedisService{}
	redisService.RedisClient = redisClient
	redisService.log = log

	return &redisService, nil
}

// CreateRedisService -- init redis
func CreateRedisService(log *zap.Logger, redisOpt *config.RedisOptions) (*RedisService, error) {
	redisService, err := NewRedisService(log, redisOpt)
	if err != nil {
		log.Error("Error", zap.Error(err))
		return nil, err
	}
	return redisService, nil
}

// Get - call the Get method on the RedisClient
func (rds *RedisService) Get(key string) (string, error) {
	fmt.Println("RedisService Get key", key)
	resp, err := rds.RedisClient.Get(key).Result()
	fmt.Println("RedisService Get resp", resp)
	fmt.Println("RedisService Get err", err)
	if err != nil {
		fmt.Println("rds.log", rds.log)
		rds.log.Error("Error", zap.Error(err))
	}

	return resp, err
}

// Set - Call the Set method on the Redis client
func (rds *RedisService) Set(key string, value interface{}, expiration time.Duration) error {
	fmt.Println("RedisService Set key", key)
	fmt.Println("RedisService Set key", value)
	err := rds.RedisClient.Set(key, value, 0).Err()
	fmt.Println("RedisService Set key value", err)
	if err != nil {
		rds.log.Error("Error", zap.Error(err))
		return err
	}

	return nil
}
