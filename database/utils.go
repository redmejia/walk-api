package database

import (
	"database/sql"

	"github.com/redmejia/walk"
)

// DB ... database connect to pool
type DB struct {
	DataBase *sql.DB
}

func (db *DB) GetProducts(query string) ([]walk.Products, error) {
	var products []walk.Products

	// rows, err := connection.DB.Query(query)
	rows, err := db.DataBase.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var product walk.Products
		err := rows.Scan(
			&product.ProductID,
			&product.ProName,
			&product.Price,
			&product.ProductImg,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
