package postgresql

import (
	product "CRUD-Go-Hexa-MongoDB/internal/domain/models"
	"CRUD-Go-Hexa-MongoDB/internal/ports"
	"database/sql"

	"github.com/google/uuid"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ports.IProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) FindAll() ([]product.Product, error) {
	rows, err := r.db.Query("SELECT id, name, stock FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []product.Product
	for rows.Next() {
		var p product.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Stock); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (r *ProductRepository) FindByID(id uuid.UUID) (product.Product, error) {
	var product product.Product
	err := r.db.QueryRow("SELECT id, name, stock FROM products WHERE id = $1", id).Scan(&product.ID, &product.Name, &product.Stock)
	return product, err
}

func (r *ProductRepository) Create(product product.Product) error {
	_, err := r.db.Exec("INSERT INTO products (id, name, stock) VALUES ($1, $2, $3)", product.ID, product.Name, product.Stock)
	return err
}

func (r *ProductRepository) Update(product product.Product) error {
	_, err := r.db.Exec("UPDATE products SET name = $1, stock = $2 WHERE id = $3", product.Name, product.Stock, product.ID)
	return err
}

func (r *ProductRepository) Delete(id uuid.UUID) error {
	_, err := r.db.Exec("DELETE FROM products WHERE id = $1", id)
	return err
}
