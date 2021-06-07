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

// RetriveById ...
func RetriveById(db *sql.DB, productID string) ProductInfo {
	var productInfo ProductInfo
	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
	}
	defer tx.Rollback()
	var product Product
	var size ProductSize
	var color ProductColor
	var img ProductImage
	err = db.QueryRow(`
			select 
				p.product_id, 
				p.pro_name, 
				p.price, 
				s.size_one, 
				s.size_two,
				s.size_three,
				s.size_four,
				c.color_one, 
				c.color_two,
				c.color_three,
				c.color_four,
				i.img_one_path, 
				i.img_two_path 
			from 
				products p 
			join 
				sizes s 
			on 
				p.product_id = s.product_id 
			join 
				colors c 
			on 
				c.product_id = p.product_id 
			join 
				shoes_img i 
			on 
				p.product_id = i.product_id 
			where  
				p.product_id = $1
	`, productID).Scan(
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
		&img.ImgTwo,
	)
	if err != nil {
		log.Fatal(err)
	}
	productInfo = ProductInfo{
		Product: product,
		Size:    []string{size.SizeOne, size.SizeTwo, size.SizeThree, size.SizeFour},
		Colors:  []string{color.ColorOne, color.ColorTwo, color.ColorThree, color.ColorFour},
		Image:   []string{img.ImgOne, img.ImgTwo},
	}
	err = tx.Commit()
	if err != nil {
		log.Println(err)
	}
	return productInfo
}

// Retrive ... take dbmodel, query, arguments and return slice of Products or interface you car asert to dbmodel
func Retrive(db *sql.DB, model interface{}, query string, args ...interface{}) ([]Products, interface{}, error) {
	switch v := model.(type) {
	case Products:
		var products []Products
		rows, err := db.Query(query, args...)
		if err != nil {
			return nil, nil, err
		}
		for rows.Next() {
			rows.Scan(&v.ProductID, &v.ProName, &v.Price, &v.ProductImg)
			products = append(products, v)
		}
		return products, nil, nil
	case Product:
		err := db.QueryRow(query, args...).Scan(&v.ProductID, &v.ProName, &v.Price)
		if err != nil {
			return nil, nil, err
		}
		return nil, v, nil
	case ProductSize:
		err := db.QueryRow(query, args...).Scan(&v.SizeOne, &v.SizeTwo, &v.SizeThree, &v.SizeFour)
		if err != nil {
			return nil, nil, err
		}
		return nil, v, nil
	case ProductColor:
		err := db.QueryRow(query, args...).Scan(&v.ColorOne, &v.ColorTwo, &v.ColorThree, &v.ColorFour)
		if err != nil {
			return nil, nil, err
		}
		return nil, v, nil
	case Signin:
		err := db.QueryRow(query, args...).Scan(&v.UserId, &v.Email, &v.Password)
		if err != nil {
			return nil, nil, err
		}
		return nil, v, nil
	default:
		log.Fatal("No matching type")
	}
	return nil, nil, nil
}

// RETRIVE COSTUMER ORDER.

// UPDATE ORDER

// DELETE ORDER
