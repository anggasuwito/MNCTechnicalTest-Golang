package model

import "database/sql"

func (HistoryLog) TableName() string {
	return "history_log"
}

type HistoryLog struct {
	Id           int          `json:"id"`
	AccessedPath string       `json:"accessed_path"`
	IpAddress    string       `json:"ip_address"`
	CreatedDate  sql.NullTime `json:"created_date"`
	UpdatedDate  sql.NullTime `json:"updated_date"`
}
