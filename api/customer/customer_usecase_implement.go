package customer

import (
	"TechnicalTest-Golang/api/model"
	"TechnicalTest-Golang/utils"
	"errors"
	"github.com/jinzhu/gorm"
)

func (c CustomerUsecase) Login(body model.Customer) error {
	var result model.Customer
	err := c.repository.GetCustomerByEmail(body.Email, &result)
	if err == gorm.ErrRecordNotFound {
		return errors.New("Customer with email " + body.Email + " doesn't exist")
	}
	if err != nil {
		return err
	}
	if body.Password != result.Password {
		return errors.New("Your email and password doesn't match")
	}
	accessToken := utils.GenerateToken(body.Email)
	c.redis.Set("token", accessToken, 0)
	c.redis.Set("email", body.Email, 0)
	return nil
}

func (c CustomerUsecase) Logout() {
	c.redis.Del("token")
}
