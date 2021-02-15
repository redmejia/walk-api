package dbutils

import (
	"database/sql"
	"errors"
)

var badData error = errors.New("A Field(s) is missing form require 5 filds.")

func NewOrder(db *sql.DB, proid uint8, name, color, size string, total float32) (sql.Result, error) {
	if proid == 0 || name == "" || color == "" || size == "" || total == 0.0 {
		return nil, badData
	} else {
		query := `INSERT INTO orders (pro_id, name, color, size, total) VALUES ($1, $2, $3, $4, $5)`
		order, err := db.Prepare(query)
		if err != nil {
			return nil, err
		}
		result, err := order.Exec(proid, name, color, size, total)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
}

func Retrive(db *sql.DB, query string) ([]Product, error) {
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	var products []Product
	for rows.Next() {
		var p Product
		rows.Scan(&p.ProID, &p.Name, &p.Color, &p.Size, &p.Price)
		products = append(products, p)
	}
	return products, nil
}

// RETRIVE COSTUMER ORDER.

// UPDATE ORDER

// DELETE ORDER
