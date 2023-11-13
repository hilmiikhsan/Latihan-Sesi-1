package product

import (
	"context"
	"database/sql"
)

type PostgresNativeRepository struct {
	db *sql.DB
}

func NewPostgresNativeRepository(db *sql.DB) PostgresNativeRepository {
	return PostgresNativeRepository{
		db: db,
	}
}

func (p PostgresNativeRepository) Create(ctx context.Context, model Product) (err error) {
	query := "INSERT INTO products (name, category, price, stock) VALUES ($1, $2, $3, $4)"

	stmt, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, model.Name, model.Category, model.Price, model.Stock)
	if err != nil {
		return
	}

	return
}

func (p PostgresNativeRepository) GetAll(ctx context.Context) (model []Product, err error) {
	query := "SELECT id, name, category, price, stock FROM products"

	rows, err := p.db.QueryContext(ctx, query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var each = Product{}
		err = rows.Scan(&each.Id, &each.Name, &each.Category, &each.Price, &each.Stock)
		if err != nil {
			return
		}

		model = append(model, each)
	}

	return
}

func (p PostgresNativeRepository) GetById(ctx context.Context, id int) (model Product, err error) {
	query := "SELECT id, name, category, price, stock FROM products WHERE id = $1"

	err = p.db.QueryRowContext(ctx, query, id).Scan(&model.Id, &model.Name, &model.Category, &model.Price, &model.Stock)
	if err != nil {
		return
	}

	return
}

func (p PostgresNativeRepository) Update(ctx context.Context, id int, model Product) (err error) {
	query := "UPDATE products SET name = $1, category = $2, price = $3, stock = $4 WHERE id = $5"

	stmt, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, model.Name, model.Category, model.Price, model.Stock, id)
	if err != nil {
		return
	}

	return
}

func (p PostgresNativeRepository) Delete(ctx context.Context, id int) (err error) {
	query := "DELETE FROM products WHERE id = $1"

	stmt, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return
	}

	return
}
