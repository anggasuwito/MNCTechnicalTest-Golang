package transaction

import (
	"TechnicalTest-Golang/api/model"
	"TechnicalTest-Golang/config"
	"TechnicalTest-Golang/utils"
	"bytes"
	"database/sql"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func initTransactionController() *gin.Engine {
	r := gin.Default()
	apiGroup := r.Group("/api")
	customerGroup := apiGroup.Group("/transaction")
	db := config.MockDatabase()
	redisConn := config.MockRedis()
	repository := NewTransactionRepository(db)
	usecase := NewTransactionUsecase(repository, redisConn)
	route := TransactionController{Usecase: usecase}
	accessToken := utils.GenerateToken("admin@gmail.com")
	utils.SetRedisValue(redisConn, "token", accessToken)
	route.Transaction(customerGroup)
	return r
}

func TestTransactionController_Payment(t *testing.T) {
	t.Run("test transaction", func(t *testing.T) {
		r := initTransactionController()
		value := model.Payment{
			Email:         "admin@gmail.com",
			BankName:      "Bank syariah",
			AccountNumber: 999001122,
			Amount:        1000000,
			CreatedDate: sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
			UpdatedDate: sql.NullTime{
				Time:  time.Time{},
				Valid: false,
			},
		}
		body, _ := json.Marshal(value)
		request, _ := http.NewRequest("POST", "/api/transaction/payment", bytes.NewReader(body))
		response := httptest.NewRecorder()
		r.ServeHTTP(response, request)
		assert.Equal(t, http.StatusOK, response.Code, "transaction success payment")
	})
}
