package product

import (
	"context"
	"log"
)

type Repository interface {
	Create(ctx context.Context, req Product) (err error)
	GetAll(ctx context.Context) (res []Product, err error)
	GetById(ctx context.Context, id int) (res Product, err error)
	Update(ctx context.Context, id int, req Product) (err error)
	Delete(ctx context.Context, id int) (err error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) CreateProduct(ctx context.Context, req Product) (err error) {
	if err = req.Validate(); err != nil {
		log.Println("error validate product", err.Error())
		return
	}

	if err = s.repo.Create(ctx, req); err != nil {
		log.Println("error create product", err.Error())
		return
	}

	return
}

func (s Service) GetAllProduct(ctx context.Context) (res []Product, err error) {
	if res, err = s.repo.GetAll(ctx); err != nil {
		log.Println("error get all product", err.Error())
		return
	}

	return
}

func (s Service) GetProductById(ctx context.Context, id int) (res Product, err error) {
	if res, err = s.repo.GetById(ctx, id); err != nil {
		log.Println("error get product by id", err.Error())
		return
	}

	return
}

func (s Service) UpdateProduct(ctx context.Context, id int, req Product) (err error) {
	if err = req.Validate(); err != nil {
		log.Println("error validate product", err.Error())
		return
	}

	_, err = s.GetProductById(ctx, id)
	if err != nil {
		log.Println("error get product by id", err.Error())
		return
	}

	if err = s.repo.Update(ctx, id, req); err != nil {
		log.Println("error update product", err.Error())
		return
	}

	return
}

func (s Service) DeleteProduct(ctx context.Context, id int) (err error) {
	_, err = s.GetProductById(ctx, id)
	if err != nil {
		log.Println("error get product by id", err.Error())
		return
	}

	if err = s.repo.Delete(ctx, id); err != nil {
		log.Println("error delete product", err.Error())
		return
	}

	return
}
