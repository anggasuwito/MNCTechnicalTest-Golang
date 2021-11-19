package model

import (
	"database/sql"
)

func (Customer) TableName() string {
	return "customer"
}

type Customer struct {
	Id          int          `json:"id"`
	Email       string       `json:"email"`
	Password    string       `json:"password"`
	CreatedDate sql.NullTime `json:"created_date"`
	UpdatedDate sql.NullTime `json:"updated_date"`
}
