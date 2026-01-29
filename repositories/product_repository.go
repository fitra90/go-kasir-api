package repositories

import (
	"database/sql"
	"errors"
	"testgo/models"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo *ProductRepository) GetAll() ([]models.Product, error) {
	rows, err := repo.db.Query("SELECT id, name, price, stock FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]models.Product, 0)
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Stock)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (repo *ProductRepository) GetByID(id int) (*models.Product, error) {
	row := repo.db.QueryRow("SELECT id, name, price, stock FROM products WHERE id = ?", id)
	var product models.Product
	err := row.Scan(&product.ID, &product.Name, &product.Price, &product.Stock)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	return &product, nil
}

func (repo *ProductRepository) Create(product *models.Product) error {

	query := "INSERT INTO products (name, price, stock) VALUES ($1, $2, $3) RETURNING id"
	err := repo.db.QueryRow(query, product.Name, product.Price, product.Stock).Scan(&product.ID)
	if err != nil {
		return err
	}

	return nil
}

func (repo *ProductRepository) Update(product *models.Product) error {
	query := "UPDATE products SET name = $1, price = $2, stock = $3 WHERE id = $4"
	result, err := repo.db.Exec(query, product.Name, product.Price, product.Stock, product.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("product not found")
	}

	return nil
}

func (repo *ProductRepository) Delete(id int) error {
	query := "DELETE FROM products WHERE id = $1"
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("product not found")
	}

	return err
}
