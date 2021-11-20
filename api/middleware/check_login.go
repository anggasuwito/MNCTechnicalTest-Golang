package middleware

import (
	"TechnicalTest-Golang/api/model"
	"TechnicalTest-Golang/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"net/http"
)

func CheckLogin(redisConn *redis.Client) gin.HandlerFunc {
	return func(context *gin.Context) {
		token, err := utils.GetRedisValue(redisConn,"token")
		if err == redis.Nil || token == "" {
			context.AbortWithStatusJSON(http.StatusUnauthorized, model.Response{Message: "You are unauthorized, you need to be login first"})
		}
		context.Next()
	}
}
