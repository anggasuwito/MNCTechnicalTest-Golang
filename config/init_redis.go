package config

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"os"
)

func InitRedis() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:         os.Getenv("REDIS_URL"),
	})

	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Panic(err)
		fmt.Print(err.Error())
	}
	return redisClient
}
