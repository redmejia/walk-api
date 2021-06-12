package dbutils

import (
	"database/sql"
	"errors"
	"log"
)

func NewOrder(db *sql.DB, proid uint8, name, color, size string, total float32) (sql.Result, error) {
	if proid == 0 || name == "" || color == "" || size == "" || total == 0.0 {
		return nil, errors.New("A Field(s) is missing form require 5 filds.")
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

// RetriveById ... retrive product by id for promotion and no promo product
func RetriveById(db *sql.DB, query, productID string) ProductInfo {
	var productInfo ProductInfo
	var product Product
	var size ProductSize
	var color ProductColor
	var img ProductImage
	err := db.QueryRow(
		query, productID).Scan(
		&product.ProductID,
		&product.ProName,
		&product.Price,
		&size.SizeOne,
		&size.SizeTwo,
		&size.SizeThree,
		&size.SizeFour,
		&color.ColorOne,
		&color.ColorTwo,
		&color.ColorThree,
		&color.ColorFour,
		&img.ImgOne,
		&img.ImgTwo)
	if err != nil {
		log.Fatal(err)
	}
	productInfo = ProductInfo{
		Product: product,
		Size:    []string{size.SizeOne, size.SizeTwo, size.SizeThree, size.SizeFour},
		Colors:  []string{color.ColorOne, color.ColorTwo, color.ColorThree, color.ColorFour},
		Image:   []string{img.ImgOne, img.ImgTwo},
	}
	return productInfo
}

// Retrive ... take dbmodel, query, arguments and return slice of Products or interface you car asert to dbmodel
func Retrive(db *sql.DB, dbmodel interface{}, query string, args ...interface{}) ([]interface{}, error) {
	var data []interface{}
	switch v := dbmodel.(type) {
	case Products:
		rows, err := db.Query(query, args...)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			err := rows.Scan(&v.ProductID, &v.ProName, &v.Price, &v.ProductImg)
			if err == sql.ErrNoRows {
				return nil, err
			}
			data = append(data, v)
		}
	case Product:
		err := db.QueryRow(query, args...).Scan(&v.ProductID, &v.ProName, &v.Price)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, err
			}
		}
		p := Product{
			ProductID: v.ProductID,
			ProName:   v.ProName,
			Price:     v.Price,
		}
		data = append(data, p)
	case ProductSize:
		err := db.QueryRow(query, args...).Scan(&v.SizeOne, &v.SizeTwo, &v.SizeThree, &v.SizeFour)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, err
			}
		}
		pS := ProductSize{
			SizeOne:   v.SizeOne,
			SizeTwo:   v.SizeTwo,
			SizeThree: v.SizeThree,
			SizeFour:  v.SizeFour,
		}
		data = append(data, pS)
	case ProductColor:
		err := db.QueryRow(query, args...).Scan(&v.ColorOne, &v.ColorTwo, &v.ColorThree, &v.ColorFour)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, err
			}
		}
		pC := ProductColor{
			ColorOne:   v.ColorOne,
			ColorTwo:   v.ColorTwo,
			ColorThree: v.ColorThree,
			ColorFour:  v.ColorFour,
		}
		data = append(data, pC)
	case Signin:
		err := db.QueryRow(query, args...).Scan(&v.UserId, &v.Email, &v.Password)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, err
			}
		}
		s := Signin{
			UserId:   v.UserId,
			Email:    v.Email,
			Password: v.Password,
		}
		data = append(data, s)
	default:
		log.Fatal("No matching type")
	}
	return data, nil
}

// RETRIVE COSTUMER ORDER.

// UPDATE ORDER

// DELETE ORDER
