package customer

import (
	"TechnicalTest-Golang/api/model"
	"github.com/jinzhu/gorm"
)

type ICustomerRepository interface {
	GetCustomerByEmail(email string, result *model.Customer) error
}

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) ICustomerRepository {
	return &CustomerRepository{db: db}
}
