package repositories

import (
	"database/sql"
	"fmt"
	"testgo/models"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (repo *TransactionRepository) CreateTransaction(items []models.CheckoutItem) (*models.Transaction, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return nil, err
	}
	// Rollback otomatis jika terjadi error sebelum Commit
	defer tx.Rollback()

	var totalAmount int
	details := make([]models.TransactionDetails, 0, len(items))

	for _, item := range items {
		var productPrice, stock int
		var productName string

		// 1. Tambahkan FOR UPDATE untuk mengunci baris produk (mencegah race condition)
		err := tx.QueryRow("SELECT name, price, stock FROM products WHERE id = $1 FOR UPDATE", item.ProductID).
			Scan(&productName, &productPrice, &stock)

		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product id %d not found", item.ProductID)
		}
		if err != nil {
			return nil, err
		}

		// 2. Cek stok secara manual 
		if stock < item.Quantity {
			return nil, fmt.Errorf("stok produk '%s' tidak mencukupi", productName)
		}

		subtotal := productPrice * item.Quantity
		totalAmount += subtotal

		// 3. Update stok
		_, err = tx.Exec("UPDATE products SET stock = stock - $1 WHERE id = $2", item.Quantity, item.ProductID)
		if err != nil {
			return nil, err
		}

		details = append(details, models.TransactionDetails{
			ProductID:   item.ProductID,
			ProductName: productName,
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
		})
	}

	// 4. Insert Header
	var transactionId int
	err = tx.QueryRow("INSERT INTO transactions (total_amount) VALUES ($1) RETURNING id", totalAmount).Scan(&transactionId)
	if err != nil {
		return nil, err
	}

	// 5. Bulk Insert untuk Details (Lebih efisien daripada loop Exec)
	if len(details) > 0 {
		query := "INSERT INTO transaction_details (transaction_id, product_id, product_name, quantity, subtotal) VALUES "
		values := []interface{}{}
		for i, d := range details {
			p := i * 5
			query += fmt.Sprintf("($%d, $%d, $%d, $%d, $%d),", p+1, p+2, p+3, p+4, p+5)
			values = append(values, transactionId, d.ProductID, d.ProductName, d.Quantity, d.Subtotal)
		}
		query = query[:len(query)-1] 

		if _, err := tx.Exec(query, values...); err != nil {
			return nil, err
		}
	}

	// 6. Commit seluruh rangkaian transaksi
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &models.Transaction{
		ID:          transactionId,
		TotalAmount: totalAmount,
		Details:     details,
	}, nil
}

// func (repo *TransactionRepository) GetTransaction(id int) (*models.CheckoutItem, error) {
// 	return nil, nil
// }
