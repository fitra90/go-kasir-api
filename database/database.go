package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitDB(connectionString string) (*sql.DB, error) {

	// open db
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// test connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	log.Println("Database connection established")

	// Auto-migrate tables
	err = migrate(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func migrate(db *sql.DB) error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS transactions (
			id SERIAL PRIMARY KEY,
			total_amount INTEGER NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS transaction_details (
			id SERIAL PRIMARY KEY,
			transaction_id INTEGER NOT NULL REFERENCES transactions(id),
			product_id INTEGER NOT NULL,
			product_name TEXT NOT NULL,
			quantity INTEGER NOT NULL,
			subtotal INTEGER NOT NULL
		)`,
	}

	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			return err
		}
	}
	return nil

}
