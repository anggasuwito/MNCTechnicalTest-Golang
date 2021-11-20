package middleware

import (
	"TechnicalTest-Golang/config"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateHistoryLog(t *testing.T) {
	t.Run("test use middleware history log", func(t *testing.T) {
		r := gin.Default()
		db := config.MockDatabase()
		var countBefore int
		var countAfter int
		r.Use(CreateHistoryLog(db))
		r.GET("/use/createhistorylog")

		db.Table("history_log").Select("*").Count(&countBefore)
		request, _ := http.NewRequest("GET", "/use/createhistorylog", nil)
		response := httptest.NewRecorder()
		r.ServeHTTP(response, request)
		db.Table("history_log").Select("*").Count(&countAfter)
		assert.Greater(t, countAfter, countBefore, "create history log added new record")
	})
}
