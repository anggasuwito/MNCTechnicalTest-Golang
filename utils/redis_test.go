package utils

import (
	"TechnicalTest-Golang/config"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetRedisValue(t *testing.T) {
	t.Run("test redis", func(t *testing.T) {
		redisConn := config.MockRedis()
		SetRedisValue(redisConn, "testKeyName", "testValue")
		value, _ := redisConn.Get("testKeyName").Result()
		assert.Equal(t, value, "testValue", "redis success set value")
	})
}

func TestGetRedisValue(t *testing.T) {
	t.Run("test redis", func(t *testing.T) {
		redisConn := config.MockRedis()
		redisConn.Set("testKeyName", "testValue", 0)
		value, _ := GetRedisValue(redisConn, "testKeyName")
		assert.Equal(t, value, "testValue", "redis success get value")
	})
}

func TestDeleteRedisValue(t *testing.T) {
	t.Run("test redis", func(t *testing.T) {
		redisConn := config.MockRedis()
		redisConn.Set("testKeyName", "testValue", 0)
		DeleteRedisValue(redisConn, "testKeyName")
		_, err := redisConn.Get("testKeyName").Result()
		assert.Equal(t, err, redis.Nil, "redis success delete value")
	})
}
