package middleware

import (
	"TechnicalTest-Golang/config"
	"TechnicalTest-Golang/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckLogin(t *testing.T) {
	t.Run("test use middleware check login", func(t *testing.T) {
		r := gin.Default()
		redisConn := config.MockRedis()
		r.Use(CheckLogin(redisConn))
		r.GET("/use/checklogin")

		request, _ := http.NewRequest("GET", "/use/checklogin", nil)
		response := httptest.NewRecorder()
		r.ServeHTTP(response, request)
		assert.Equal(t, http.StatusUnauthorized, response.Code, "check login not pass")

		utils.SetRedisValue(redisConn, "token", "accessToken")
		request2, _ := http.NewRequest("GET", "/use/checklogin", nil)
		response2 := httptest.NewRecorder()
		r.ServeHTTP(response2, request2)
		assert.Equal(t, http.StatusOK, response2.Code, "check login pass")
	})
}
