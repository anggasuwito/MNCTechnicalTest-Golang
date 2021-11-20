package config

import (
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis"
)

func MockRedis() *redis.Client {
	mockServer, _ := miniredis.Run()
	redisClient := redis.NewClient(&redis.Options{Addr: mockServer.Addr()})
	return redisClient
}
