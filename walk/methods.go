package walk

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/redmejia/connection"
)

// GetProducts ... Retrive categories product
func (p *Products) GetProducts(query string) ([]Products, error) {
	var products []Products
	rows, err := connection.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var product Products
		err := rows.Scan(&product.ProductID, &product.ProName, &product.Price, &product.ProductImg)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

// GetProductById ... Retrive product by id
func (p *ProductInfo) GetProductById(query, productID string) ProductInfo {
	var product Product
	var size ProductSize
	var color ProductColor
	var img ProductImage

	row := connection.DB.QueryRow(query, productID)

	err := row.Scan(
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

	productInfo := ProductInfo{
		Product: product,
		Size:    []string{size.SizeOne, size.SizeTwo, size.SizeThree, size.SizeFour},
		Colors:  []string{color.ColorOne, color.ColorTwo, color.ColorThree, color.ColorFour},
		Image:   []string{img.ImgOne, img.ImgTwo},
	}

	return productInfo
}
