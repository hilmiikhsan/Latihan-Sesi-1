package database

import (
	"fmt"
	"latihan-bottcamp/crud_gorm_sqlx/app/product"
	"latihan-bottcamp/crud_gorm_sqlx/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectGORMPostgres(cfg config.DB) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name,
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	db.Debug().AutoMigrate(product.Product{})

	return
}
