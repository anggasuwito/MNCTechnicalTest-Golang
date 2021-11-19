package transaction

import "TechnicalTest-Golang/api/model"

func (t TransactionRepository) GetCustomerByEmail(email string, result *model.Customer) error {
	err := t.db.Debug().Table("customer").Select("*").Where("email = ?", email).Scan(&result).Error
	if err != nil {
		return err
	}
	return nil
}

func (t TransactionRepository) CreateNewPayment(body model.Payment) error {
	err := t.db.Debug().Table("payment").Create(&body).Error
	if err != nil {
		return err
	}
	return nil
}
