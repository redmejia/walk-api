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

// NewOrder ...
func (c *ClientOrder) NewOrder(status *PurchaseStatus) {
	db := connection.DB

	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
	}

	defer tx.Rollback()

	var purchaseID int
	row := tx.QueryRow(
		`INSERT INTO client_info(
				user_id,
				first_name,
				last_name,
				email,
				address,
				state,
				zip,
				name_on_card,
				card_number,
				cv_number)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
			RETURNING purchase_id`,
		c.Client.UserId,
		c.Client.FirstName,
		c.Client.LastName,
		c.Client.Email,
		c.Client.Address,
		c.Client.State,
		c.Client.Zip,
		c.Client.NameOnCard,
		c.Client.CardNumber,
		c.Client.CvNumber,
	)

	err = row.Scan(&purchaseID)

	if err != nil {
		log.Println(err)
	}

	for _, v := range c.Items {
		_, err = tx.Exec(`
			INSERT INTO client_order(
				purchase_id,
				user_id,
				product_id,
				pro_name,
				color,
				size,
				qty,
				price)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
			purchaseID,
			c.Client.UserId,
			v.ProductId,
			v.ProName,
			v.Color,
			v.Size,
			v.Qty,
			v.Price)
		if err != nil {
			log.Println("client_order ", err)
		}
	}

	_, err = tx.Exec(`
		INSERT INTO client_order_total(purchase_id, user_id, total)
		VALUES ($1, $2, $3)`, purchaseID, c.Client.UserId, c.Total)

	if err != nil {
		log.Println(err)
	}

	_, err = tx.Exec(`INSERT INTO  client_purchase_status(purchase_id, user_id, purchase_status, purchase_code)
			VALUES ($1, $2, $3, $4)`, purchaseID, c.Client.UserId, status.Status, status.TransactionCode)

	if err != nil {
		log.Println(err)
	}

	err = tx.Commit()

	if err != nil {
		log.Println(err)
	}
}
