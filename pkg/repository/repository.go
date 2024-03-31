package repository

import (
	"pro_pay/tools/logger"

	"pro_pay/pkg/repository/atmost"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repository struct {
	AtmostRepo *atmost.Atmost
}

func NewRepository(db *sqlx.DB, loggers *logger.Logger) *Repository {
	return &Repository{

		AtmostRepo: atmost.NewPaymentRepo(db, loggers),
	}
}
