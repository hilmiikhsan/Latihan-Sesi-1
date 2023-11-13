package main

import (
	"latihan-bottcamp/crud_gorm_sqlx/app/product"
	"latihan-bottcamp/crud_gorm_sqlx/config"
	"latihan-bottcamp/crud_gorm_sqlx/pkg/database"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	router := fiber.New(fiber.Config{
		AppName: "Products Services",
		Prefork: true,
	})

	err := config.LoadConfig("./config/config.yaml")
	if err != nil {
		log.Println("error when try to LoadConfig with error :", err.Error())
	}

	// dbGorm, err := database.ConnectGORMPostgres(config.Cfg.DB)
	// if err != nil {
	// 	panic(err)
	// }

	// dbSqlx, err := database.ConnectSQLXPostgres(config.Cfg.DB)
	// if err != nil {
	// 	panic(err)
	// }

	db, err := database.ConnectNativePostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	product.RegisterServiceProduct(router, product.DB{Db: db})

	router.Listen(config.Cfg.App.Port)
}
