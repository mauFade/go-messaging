package repository

import (
	"database/sql"

	"github.com/mauFade/go-messaging/internal/entity"
)

type ProductRepository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) Create(product *entity.Product) error {
	_, err := r.DB.Exec("INSERT INTO products (id, name, price) VALUES(?, ?, ?)",
		product.ID, product.Name, product.Price)

	if err != nil {
		return err
	}

	return nil
}

func (r *ProductRepository) Find() ([]*entity.Product, error) {
	rows, err := r.DB.Query("SELECT id, name, price, products FROM products")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []*entity.Product

	for rows.Next() {
		var product entity.Product

		err = rows.Scan(&product.ID, &product.Name, &product.Price)

		if err != nil {
			return nil, err
		}

		products = append(products, &product)
	}

	return products, nil
}
