package repositories

import (
	"database/sql"
	"testgo/models"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (repo *ReportRepository) GetReport(startDate, endDate string) (*models.Report, error) {
	var report models.Report
	var args []interface{}

	// 1. Get Total Revenue and Total Transaction
	query := `
		SELECT 
			COALESCE(SUM(total_amount), 0), 
			COUNT(id)
		FROM transactions 
	`
	if startDate != "" && endDate != "" {
		query += " WHERE created_at BETWEEN $1 AND $2"
		args = append(args, startDate, endDate)
	}

	err := repo.db.QueryRow(query, args...).Scan(&report.TotalRevenue, &report.TotalTransaksi)
	if err != nil {
		return nil, err
	}

	// 2. Get Best Selling Product
	queryProdukTerlaris := `
		SELECT 
			p.name, 
			COALESCE(SUM(td.quantity), 0) as total_qty
		FROM transaction_details td
		JOIN transactions t ON td.transaction_id = t.id
		JOIN products p ON td.product_id = p.id
	`

	// Reset args for the second query
	args = []interface{}{}
	if startDate != "" && endDate != "" {
		queryProdukTerlaris += " WHERE t.created_at BETWEEN $1 AND $2"
		args = append(args, startDate, endDate)
	}

	queryProdukTerlaris += `
		GROUP BY p.id, p.name
		ORDER BY total_qty DESC
		LIMIT 1
	`

	err = repo.db.QueryRow(queryProdukTerlaris, args...).Scan(&report.ProdukTerlaris.Nama, &report.ProdukTerlaris.QtyTerjual)
	if err != nil {
		if err == sql.ErrNoRows {
			// No products sold in this period (or at all), which is valid.
			report.ProdukTerlaris = models.ProdukTerlaris{}
		} else {
			return nil, err
		}
	}

	return &report, nil
}
