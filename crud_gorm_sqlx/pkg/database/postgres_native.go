package database

import (
	"database/sql"
	"fmt"
	"latihan-bottcamp/crud_gorm_sqlx/config"
)

func ConnectNativePostgres(cfg config.DB) (db *sql.DB, err error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name,
	)

	db, err = sql.Open("postgres", dsn)
	if err != nil {
		return
	}

	if err = db.Ping(); err != nil {
		return
	}

	return
}
