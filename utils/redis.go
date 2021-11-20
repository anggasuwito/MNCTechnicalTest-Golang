package utils

import "github.com/go-redis/redis"

func SetRedisValue(redisConn *redis.Client, keyName string, value string) {
	redisConn.Set(keyName, value, 0)
}

func GetRedisValue(redisConn *redis.Client, keyName string) (string, error) {
	result, err := redisConn.Get(keyName).Result()
	if err != nil {
		return result, err
	}
	return result, nil
}

func DeleteRedisValue(redisConn *redis.Client, keyName string) {
	redisConn.Del(keyName)
}
