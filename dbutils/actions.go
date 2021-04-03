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

func Retrive(db *sql.DB, model interface{}, query string, targets ...interface{}) ([]interface{}, error) {
	var data []interface{}
	switch v := model.(type) {
	case Product:
		rows, err := db.Query(query, targets...)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			// var p v
			rows.Scan(&v.ProID, &v.Name, &v.Color, &v.Size, &v.Price)
			data = append(data, v)
		}
	case SigninForm:
		rows, err := db.Query(query, targets...)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			// var p v
			rows.Scan(&v.Email, &v.Password)
			data = append(data, v)
		}
	}
	return data, nil
}

// RETRIVE COSTUMER ORDER.

// UPDATE ORDER

// DELETE ORDER
