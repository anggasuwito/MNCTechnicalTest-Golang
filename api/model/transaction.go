package model

import (
	"database/sql"
)

func (Payment) TableName() string {
	return "payment"
}

type Payment struct {
	Id            int          `json:"id"`
	Email         string       `json:"email"`
	BankName      string       `json:"bank_name"`
	AccountNumber int          `json:"account_number"`
	Amount        int          `json:"amount"`
	CreatedDate   sql.NullTime `json:"created_date"`
	UpdatedDate   sql.NullTime `json:"updated_date"`
}
