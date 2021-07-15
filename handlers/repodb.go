package handlers

import (
	"database/sql"

	"github.com/redmejia/walk"
)

type DBRepo struct {
	Conn *sql.DB
}

func (db *DBRepo) GetProducts(query string) ([]walk.Products, error) {
	var products []walk.Products

	// rows, err := connection.DB.Query(query)
	rows, err := db.Conn.Query(query)

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

type DataBaseRepo interface {
	GetProducts(query string) ([]walk.Products, error)
}
