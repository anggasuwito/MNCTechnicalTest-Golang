package transaction

import (
	"TechnicalTest-Golang/api/model"
	"github.com/jinzhu/gorm"
)

type ITransactionRepository interface {
	GetCustomerByEmail(email string, result *model.Customer) error
	CreateNewPayment(body model.Payment) error
}

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) ITransactionRepository {
	return &TransactionRepository{db: db}
}
