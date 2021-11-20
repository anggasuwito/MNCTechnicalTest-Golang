package transaction

import (
	"TechnicalTest-Golang/api/model"
	"TechnicalTest-Golang/utils"
	"database/sql"
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

func (t TransactionUsecase) Payment(body model.Payment) (model.Payment, error) {
	var result model.Payment
	var resultCustomer model.Customer
	err := t.repository.GetCustomerByEmail(body.Email, &resultCustomer)
	if err == gorm.ErrRecordNotFound {
		return result, errors.New("Account with email " + body.Email + " is not been registered")
	}
	if err != nil {
		return result, err
	}
	token, err := utils.GetRedisValue(t.redis,"token")
	if err != nil {
		return result, err
	}
	isValidatedToken, email := utils.ValidateToken(token)
	if !isValidatedToken {
		return result, errors.New("You token has expired you need to be relogin")
	}
	if body.Email != email {
		return result, errors.New("You cant do payment with different logged in email")
	}
	body.CreatedDate = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	body.UpdatedDate = sql.NullTime{
		Time:  time.Time{},
		Valid: false,
	}
	err = t.repository.CreateNewPayment(body)
	result = body
	return result, nil
}
