package config

import (
	"TechnicalTest-Golang/api/model"
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func MockDatabase() *gorm.DB {
	sqlite, _ := sql.Open("sqlite3", ":memory:")

	db, _ := gorm.Open("sqlite", sqlite)
	db.AutoMigrate(&model.Customer{})
	db.AutoMigrate(&model.Payment{})
	db.AutoMigrate(&model.HistoryLog{})
	db.Exec("insert into customer (email,password) values ('admin@gmail.com','123')")
	db.Exec("insert into customer (email,password) values ('admintest@gmail.com','123')")

	return db
}
