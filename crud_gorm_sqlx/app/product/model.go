package product

import "errors"

var (
	ErrEmptyName     = errors.New("nama tidak boleh kosoong")
	ErrEmptyCategory = errors.New("category tidak boleh kosoong")
	ErrEmptyPrice    = errors.New("price tidak boleh kosoong")
	ErrEmptyStock    = errors.New("stock tidak boleh kosoong")
)

type Product struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int    `json:"price"`
	Stock    int    `json:"stock"`
}

func (p Product) Validate() (err error) {
	if p.Name == "" {
		return ErrEmptyName
	}

	if p.Category == "" {
		return ErrEmptyCategory
	}

	if p.Price == 0 {
		return ErrEmptyPrice
	}

	if p.Stock == 0 {
		return ErrEmptyStock
	}

	return
}
