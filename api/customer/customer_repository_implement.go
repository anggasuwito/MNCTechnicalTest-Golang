package customer

import "TechnicalTest-Golang/api/model"

func (c CustomerRepository) GetCustomerByEmail(email string, result *model.Customer) error {
	err := c.db.Debug().Table("customer").Select("*").Where("email = ?", email).Scan(&result).Error
	if err != nil {
		return err
	}
	return nil
}
