package repository

import (
	"pro_pay/config"
	"pro_pay/tools/logger"
	"database/sql"
	"fmt"

	migrate "github.com/rubenv/sql-migrate"
)

func psqlInitURL(cfg config.Database) string {
	// URL for Migration
	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.UserName,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.SSLMode,
	)
	return url
}
func MigratePsql(cfg config.Database, loggers *logger.Logger, up bool) error {
	migrations := &migrate.FileMigrationSource{
		Dir: "./schema",
	}
	db, err := sql.Open("postgres", psqlInitURL(cfg))
	if err != nil {
		loggers.Error("error in creating migrations: ", err.Error())
	}
	defer db.Close()
	migrateState := migrate.Up
	if !up {
		migrateState = migrate.Down
	}
	n, err := migrate.Exec(db, "postgres", migrations, migrateState)
	loggers.Info(n)
	if err != nil {
		loggers.Error("error in creating migrations: ", err.Error())
		return err
	}
	if n > 0 {
		loggers.Info("migrations applied: ", n)
	}
	return nil
}
