package api

import (
	"TechnicalTest-Golang/api/customer"
	"TechnicalTest-Golang/api/middleware"
	"TechnicalTest-Golang/api/transaction"
	"TechnicalTest-Golang/config"
	"github.com/gin-gonic/gin"
	"os"
)

func Init() {
	r := gin.Default()
	db := config.InitDb()
	redis := config.InitRedis()

	r.GET("", func(context *gin.Context) {
		return
	})

	apiGroup := r.Group("/api")
	apiGroup.Use(middleware.CreateHistoryLog(db))

	customerGroup := apiGroup.Group("/customer")
	customerRepository := customer.NewCustomerRepository(db)
	customerUsecase := customer.NewCustomerUsecase(customerRepository, redis)
	customerRoute := customer.CustomerController{Usecase: customerUsecase}
	customerRoute.Customer(customerGroup)

	transactionGroup := apiGroup.Group("/transaction")
	transactionRepository := transaction.NewTransactionRepository(db)
	transactionUsecase := transaction.NewTransactionUsecase(transactionRepository, redis)
	transactionRoute := transaction.TransactionController{Usecase: transactionUsecase}
	transactionGroup.Use(middleware.CheckLogin(redis))
	transactionRoute.Transaction(transactionGroup)

	r.Run(os.Getenv("PORT"))
}
