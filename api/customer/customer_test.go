package customer

import (
	"TechnicalTest-Golang/api/model"
	"TechnicalTest-Golang/config"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func initCustomerController() *gin.Engine {
	r := gin.Default()
	apiGroup := r.Group("/api")
	customerGroup := apiGroup.Group("/customer")
	db := config.MockDatabase()
	redisConn := config.MockRedis()
	repository := NewCustomerRepository(db)
	usecase := NewCustomerUsecase(repository, redisConn)
	route := CustomerController{Usecase: usecase}
	route.Customer(customerGroup)
	return r
}

func TestCustomerController_Login(t *testing.T) {
	t.Run("test customer", func(t *testing.T) {
		r := initCustomerController()
		value := model.Customer{
			Email:    "admin@gmail.com",
			Password: "123",
		}
		body, _ := json.Marshal(value)
		request, _ := http.NewRequest("POST", "/api/customer/login", bytes.NewReader(body))
		response := httptest.NewRecorder()
		r.ServeHTTP(response, request)
		assert.Equal(t, http.StatusOK, response.Code,"customer login success response")
	})
}

func TestCustomerController_Logout(t *testing.T) {
	t.Run("test customer", func(t *testing.T) {
		r := initCustomerController()
		request, _ := http.NewRequest("POST", "/api/customer/logout", nil)
		response := httptest.NewRecorder()
		r.ServeHTTP(response, request)
		assert.Equal(t, http.StatusOK, response.Code,"customer logout success response")
	})
}
