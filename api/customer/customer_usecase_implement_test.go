package customer

import (
	"TechnicalTest-Golang/api/model"
	"TechnicalTest-Golang/config"
	"TechnicalTest-Golang/utils"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCustomerUsecase_Login(t *testing.T) {
	t.Run("test customer usecase", func(t *testing.T) {
		value1 := model.Customer{
			Email:    "admin@gmail.com",
			Password: "123",
		}
		value2 := model.Customer{
			Email:    "adminnotexist@gmail.com",
			Password: "123",
		}
		value3 := model.Customer{
			Email:    "admin@gmail.com",
			Password: "wrongpassword",
		}
		db := config.MockDatabase()
		redisConn := config.MockRedis()
		repository := NewCustomerRepository(db)
		usecase := NewCustomerUsecase(repository, redisConn)
		err1 := usecase.Login(value1)
		err2 := usecase.Login(value2)
		err3 := usecase.Login(value3)
		assert.Nil(t, err1, "customer usecase login success")
		assert.NotNil(t, err2, "customer usecase login fail email not exist")
		assert.NotNil(t, err3, "customer usecase login fail wrong password")
	})
}

func TestCustomerUsecase_Logout(t *testing.T) {
	t.Run("test customer usecase", func(t *testing.T) {
		db := config.MockDatabase()
		redisConn := config.MockRedis()
		repository := NewCustomerRepository(db)
		usecase := NewCustomerUsecase(repository, redisConn)
		usecase.Logout()
		_, err := utils.GetRedisValue(redisConn, "token")
		assert.Equal(t, err, redis.Nil,"customer usecase logout delete stored access token")
	})
}
