package atmost

import (
	"pro_pay/model"
	"pro_pay/tools/logger"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Atmost struct {
	Transaction
}


type Transaction interface {
	CreateTransaction(in model.CreateTransaction) (id uuid.UUID, err error)
}

func NewPaymentRepo(db *sqlx.DB, loggers *logger.Logger) *Atmost {
	return &Atmost{
		Transaction: NewTransactionDB(db, loggers),
	}
}
