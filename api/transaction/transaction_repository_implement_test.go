package transaction

import (
	"TechnicalTest-Golang/api/model"
	"TechnicalTest-Golang/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransactionRepository_GetCustomerByEmail(t *testing.T) {
	t.Run("test transaction repository", func(t *testing.T) {
		var result model.Customer
		db := config.MockDatabase()
		repository := NewTransactionRepository(db)
		repository.GetCustomerByEmail("admin@gmail.com", &result)
		assert.Equal(t, result.Email, "admin@gmail.com", "transaction repository found customer by email")

		var result2 model.Customer
		repository.GetCustomerByEmail("adminnotexist@gmail.com", &result2)
		assert.Equal(t, result2.Email, "", "transaction repository not found customer by email")
	})
}

func TestTransactionRepository_CreateNewPayment(t *testing.T) {
	t.Run("test transaction repository", func(t *testing.T) {
		var body model.Payment
		var countRecordBefore int
		var countRecordAfter int
		db := config.MockDatabase()
		db.Table("payment").Select("*").Count(&countRecordBefore)
		repository := NewTransactionRepository(db)
		repository.CreateNewPayment(body)
		db.Table("payment").Select("*").Count(&countRecordAfter)
		assert.Greater(t, countRecordAfter, countRecordBefore, "transaction repository record added when create new payment")
	})
}
