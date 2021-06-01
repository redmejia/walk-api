package dbutils

import (
	"database/sql"
	"errors"
	"log"
)

var badData error = errors.New("A Field(s) is missing form require 5 filds.")

func NewOrder(db *sql.DB, proid uint8, name, color, size string, total float32) (sql.Result, error) {
	if proid == 0 || name == "" || color == "" || size == "" || total == 0.0 {
		return nil, badData
	} else {
		orderStm, err := db.Prepare(`INSERT INTO orders (pro_id, name, color, size, total) VALUES ($1, $2, $3, $4, $5)`)
		if err != nil {
			return nil, err
		}
		result, err := orderStm.Exec(proid, name, color, size, total)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
}

func Retrive(db *sql.DB, model interface{}, query string, args ...interface{}) ([]Products, interface{}, error) {
	switch v := model.(type) {
	case Products:
		var products []Products
		rows, err := db.Query(query, args...)
		if err != nil {
			return nil, nil, err
		}
		for rows.Next() {
			rows.Scan(&v.ProID, &v.ProductID, &v.ProName, &v.Price)
			products = append(products, v)
		}
		return products, nil, nil
	case Product:
		rows, err := db.Query(query, args...)
		if err != nil {
			return nil, nil, err
		}
		for rows.Next() {
			rows.Scan(&v.ProductID, &v.ProName, &v.Price)
			return nil, v, nil
		}
	case Sizes:
		rows, err := db.Query(query, args...)
		if err != nil {
			return nil, nil, err
		}
		for rows.Next() {
			rows.Scan(&v.ProductId, &v.SizeOne, &v.SizeTwo, &v.SizeThree, &v.SizeFour)
			return nil, v, nil
		}

		// err := db.QueryRow(query, args...).Scan(&v.SizeOne, &v.SizeTwo, &v.SizeThree, &v.SizeFour)
		// if err != nil {
		// 	return nil, nil, err
		// }
		// return nil, v, nil
	case Signin:
		rows, err := db.Query(query, args...)
		if err != nil {
			return nil, nil, err
		}
		for rows.Next() {
			rows.Scan(&v.UserId, &v.Email, &v.Password)
			return nil, v, nil
		}
	default:
		log.Fatal("No matching type")
	}
	return nil, nil, nil
}

// RETRIVE COSTUMER ORDER.

// UPDATE ORDER

// DELETE ORDER
