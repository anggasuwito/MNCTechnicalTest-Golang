package config

import (
	"TechnicalTest-Golang/api/model"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

func InitDb() *gorm.DB {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))
	db, err := gorm.Open("mysql", dataSource)
	db.AutoMigrate(&model.Customer{})
	db.AutoMigrate(&model.Payment{})
	db.AutoMigrate(&model.HistoryLog{})
	if err != nil {
		log.Println("FAIL CONFIG CONNECT TO DB MYSQL")
		panic(err)
		return nil
	}

	var result model.Customer
	customer1 := model.Customer{Email: "customer1@gmail.com", Password: "123"}
	customer2 := model.Customer{Email: "customer2@gmail.com", Password: "123"}
	err1 := db.Table("customer").Select("*").Where("email = ?", customer1.Email).Scan(&result).Error
	err2 := db.Table("customer").Select("*").Where("email = ?", customer2.Email).Scan(&result).Error
	if err1 == gorm.ErrRecordNotFound {
		db.Create(&customer1)
	}
	if err2 == gorm.ErrRecordNotFound {
		db.Create(&customer2)
	}

	return db
}
