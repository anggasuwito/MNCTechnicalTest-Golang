package transaction

import (
	"TechnicalTest-Golang/api/model"
	"TechnicalTest-Golang/config"
	"TechnicalTest-Golang/utils"
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTransactionUsecase_Payment(t *testing.T) {
	t.Run("test transaction usecase", func(t *testing.T) {
		value1 := model.Payment{
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
		value2 := model.Payment{
			Email:         "adminnotexist@gmail.com",
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
		value3 := model.Payment{
			Email:         "admintest@gmail.com",
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
		db := config.MockDatabase()
		redisConn := config.MockRedis()
		repository := NewTransactionRepository(db)
		usecase := NewTransactionUsecase(repository, redisConn)
		accessToken := utils.GenerateToken("admin@gmail.com")
		utils.SetRedisValue(redisConn, "token", accessToken)
		_, err := usecase.Payment(value1)
		_, err2 := usecase.Payment(value2)
		_, err3 := usecase.Payment(value3)
		assert.Nil(t, err, "transaction usecase correct data payment")
		assert.NotNil(t, err2, "transaction usecase user not registered")
		assert.NotNil(t, err3, "transaction usecase payment with different logged in account")
	})
}
