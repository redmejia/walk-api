package makeorder

import (
	"database/sql"
	"errors"
)

type Product struct {
	ProID uint8   `json:"pro_id"`
	Name  string  `json:"name"`
	Color string  `json:"color"`
	Size  string  `json:"size"`
	Total float32 `json:"total"`
}

var badData error = errors.New("A Field(s) is missing form require 5 filds.")

func newOrder(db *sql.DB, proid uint8, name, color, size string, total float32) (sql.Result, error) {
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
