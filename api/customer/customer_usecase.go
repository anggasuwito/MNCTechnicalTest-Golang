package customer

import (
	"TechnicalTest-Golang/api/model"
	"github.com/go-redis/redis"
)

type ICustomerUsecase interface {
	Login(body model.Customer) error
	Logout()
}

type CustomerUsecase struct {
	redis      *redis.Client
	repository ICustomerRepository
}

func NewCustomerUsecase(repository ICustomerRepository, redis *redis.Client) ICustomerUsecase {
	return &CustomerUsecase{repository: repository, redis: redis}
}
