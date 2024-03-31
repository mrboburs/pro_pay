package atmost

import (
	"errors"
	"pro_pay/model"
	"pro_pay/tools/logger"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

var (
	ErrorNoRowsAffected = errors.New("no rows affected")
)

type TransactionDB struct {
	db      *sqlx.DB
	loggers *logger.Logger
}

func NewTransactionDB(db *sqlx.DB, loggers *logger.Logger) *TransactionDB {
	return &TransactionDB{db: db, loggers: loggers}
}

func (repo *TransactionDB) CreateTransaction(in model.CreateTransaction) (id uuid.UUID,
	err error) {
	loggers := repo.loggers
	db := repo.db
	row, err := db.Query(
		`INSERT INTO transaction (
			amount ,
			account,
			terminal_id,
			store_id 
			) VALUES (
			$1, 
			$2,
			$3,
			$4
		) RETURNING id`,
		in.Amount, in.Account, 99, in.StoreId)
	if err != nil {
		loggers.Error(err)
		return id, err
	}
	for row.Next() {
		err = row.Scan(&id)
		if err != nil {
			loggers.Error(err)
			return id, err
		}
	}
	return id, nil
}
