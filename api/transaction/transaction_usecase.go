package transaction

import (
	"TechnicalTest-Golang/api/model"
	"github.com/go-redis/redis"
)

type ITransactionUsecase interface {
	Payment(body model.Payment) (model.Payment, error)
}

type TransactionUsecase struct {
	repository ITransactionRepository
	redis      *redis.Client
}

func NewTransactionUsecase(repository ITransactionRepository, redis *redis.Client) ITransactionUsecase {
	return &TransactionUsecase{repository: repository, redis: redis}
}
