package atmost

import (
	"pro_pay/config"
	"pro_pay/model"
	"pro_pay/pkg/repository"
	"pro_pay/pkg/store"
	"pro_pay/tools/logger"
)

type Atmost struct {
	Transaction
}

type Transaction interface {
	CreateTransaction(in model.CreateTransaction) (out *model.Response, err error)
}

func NewAtmostService(repo *repository.Repository, minio *store.Store,
	config *config.Configuration, loggers *logger.Logger) *Atmost {
	return &Atmost{

		Transaction: NewTransactionService(repo, minio, loggers),
	}
}
