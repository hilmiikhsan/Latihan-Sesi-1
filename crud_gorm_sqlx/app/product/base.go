package product

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

type DB struct {
	Db     *sql.DB
	Dbx    *sqlx.DB
	GormDB *gorm.DB
}

func RegisterServiceProduct(router fiber.Router, db DB) {
	// repo := NewPostgresGormRepository(db.GormDB)
	// repo := NewPostgresSQLXRepository(db.Dbx)
	repo := NewPostgresNativeRepository(db.Db)
	svc := NewService(repo)
	handler := NewHandler(svc)

	var productRouter = router.Group("products")
	{
		productRouter.Post("/", handler.CreateProduct)
		productRouter.Get("/", handler.GetAllProduct)
		productRouter.Get("/:id", handler.GetProductById)
		productRouter.Put("/:id", handler.UpdateProduct)
		productRouter.Delete("/:id", handler.DeleteProduct)
	}
}
