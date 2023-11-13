package product

import (
	"context"

	"gorm.io/gorm"
)

type PostgresGormRepository struct {
	db *gorm.DB
}

func NewPostgresGormRepository(db *gorm.DB) PostgresGormRepository {
	return PostgresGormRepository{
		db,
	}
}

func (p PostgresGormRepository) Create(ctx context.Context, model Product) (err error) {
	return p.db.Create(&model).Error
}

func (p PostgresGormRepository) GetAll(ctx context.Context) (model []Product, err error) {
	err = p.db.Find(&model).Error
	return
}

func (p PostgresGormRepository) GetById(ctx context.Context, id int) (model Product, err error) {
	err = p.db.First(&model, id).Error
	return
}

func (p PostgresGormRepository) Update(ctx context.Context, id int, model Product) (err error) {
	err = p.db.Model(&Product{}).Where("id = ?", id).Updates(model).Error
	return
}

func (p PostgresGormRepository) Delete(ctx context.Context, id int) (err error) {
	err = p.db.Delete(&Product{}, id).Error
	return
}
