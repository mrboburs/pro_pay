package repository

import (
	"pro_pay/config"
	"pro_pay/tools/logger"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresDB(cfg *config.Database, loggers *logger.Logger) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.UserName, cfg.Name, cfg.Password, cfg.SSLMode))
	if err != nil {
		loggers.Fatalf("failed check db configs.%v", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		loggers.Fatalf("fail ping to db %v", err)
		return nil, err
	}
	return db, nil
}
