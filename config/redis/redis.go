package redis

import (
	"context"
	"fmt"

	"github.com/payslip/config"
	"github.com/redis/go-redis/v9"
)

func NewRedisCache(config *config.Config) (*redis.Client, error) {
	opt, err := redis.ParseURL(config.RedisUrl)
	if err != nil {
		return nil, fmt.Errorf("error parse url: %v", err)
	}
	redisClient := redis.NewClient(opt)
	// Ping Redis to check connection
	_, err = redisClient.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %v", err)
	}

	return redisClient, nil
}
