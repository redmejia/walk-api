package database

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/redmejia/walk"
)

// GetProducts ... Retrive categories product
func (d *DataBase) GetProducts(query string) ([]walk.Products, error) {
	var products []walk.Products

	rows, err := d.Conn.Query(query)
	if err != nil {
		log.Println(" err ", err)
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

// GetProductById ... Retrive product by id
func (d *DataBase) GetProductById(query string, productID int) walk.ProductInfo {
	var product walk.Product
	var size walk.ProductSize
	var color walk.ProductColor
	var img walk.ProductImage

	row := d.Conn.QueryRow(query, productID)

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

	return walk.ProductInfo{
		Product: product,
		Size:    []string{size.SizeOne, size.SizeTwo, size.SizeThree, size.SizeFour},
		Colors:  []string{color.ColorOne, color.ColorTwo, color.ColorThree, color.ColorFour},
		Image:   []string{img.ImgOne, img.ImgTwo},
	}

}

// InsertNewOrder ...
func (d *DataBase) InsertNewOrder(c *walk.ClientOrder, status walk.PurchaseStatus) {
	tx, err := d.Conn.Begin()
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
		INSERT INTO 
			client_order_total(purchase_id, user_id, total)
		VALUES ($1, $2, $3)`,
		purchaseID, c.Client.UserId, c.Total,
	)

	if err != nil {
		log.Println(err)
	}

	_, err = tx.Exec(`
			INSERT INTO  
				client_purchase_status(purchase_id, user_id, purchase_status, purchase_code)
			VALUES 
				($1, $2, $3, $4)`,
		purchaseID, c.Client.UserId, status.Status, status.TransactionCode,
	)

	if err != nil {
		log.Println(err)
	}

	err = tx.Commit()

	if err != nil {
		log.Println(err)
	}
}

// GetClientPurchaseInfoByUserId ... retrive client purchase information
func (d *DataBase) GetClientPurchaseInfoByUserId(userId int) *walk.Purchase {
	var order []walk.Order
	rows, err := d.Conn.Query(`
		SELECT 	ci.purchase_id,
			ci.user_id,	
			ci.first_name,
			ci.last_name,
			ci.email,
			ci.address,
			ci.state,
			ci.zip,
			o.purchase_id,
			o.product_id,
			o.pro_name,
			o.color,
			o.size,
			o.qty,
			o.price,
			s.img_one_path,
			ct.total,
			cs.purchase_code as status_code
		FROM 
			client_info ci
		JOIN 
			client_order o ON ci.user_id = ci.user_id
		AND 
			o.purchase_id = ci.purchase_id
		JOIN 
			shoes_img s ON o.product_id = s.product_id
		JOIN 
			client_purchase_status cs ON o.user_id = cs.user_id
		AND 
			cs.purchase_id = o.purchase_id
		JOIN 
			client_order_total ct ON o.user_id = ct.user_id
		AND 
			ct.purchase_id = o.purchase_id
		WHERE 
			ci.user_id = $1 
		ORDER BY ci.purchase_id
	`, userId)

	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var pt walk.Order
		rows.Scan(
			&pt.Client.PurchaseID, &pt.Client.UserId, &pt.Client.FirstName, &pt.Client.LastName, &pt.Client.Email,
			&pt.Client.Address, &pt.Client.State, &pt.Client.Zip, &pt.Product.PurchaseID, &pt.Product.ProductId,
			&pt.Product.ProName, &pt.Product.Color, &pt.Product.Size, &pt.Product.Qty,
			&pt.Product.Price, &pt.Product.Img, &pt.Product.Total, &pt.Product.StatusCode,
		)
		order = append(order, pt)
	}
	return &walk.Purchase{
		Order: order,
	}

}

// ClientRegister ... register new user
func (d *DataBase) ClientRegister(c *walk.ClientRegister, w http.ResponseWriter) {
	tx, err := d.Conn.Begin()

	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	var userId int

	row := tx.QueryRow(`
			INSERT INTO 
				register (name, email) 
			VALUES ($1, $2) 
			RETURNING user_id`, c.Name, c.Email,
	)

	err = row.Scan(&userId)

	if err != nil {
		log.Fatal(err)
	}

	var email string

	row = tx.QueryRow(`
			INSERT INTO 
				signin (user_id, email, password) 
			VALUES ($1, $2, $3)
			RETURNING email
			`, userId, c.Email, c.Password,
	)

	err = row.Scan(&email)
	if err != nil {
		log.Fatal(err)
	}

	userName := strings.Split(email, "@")[0]

	res := walk.Message{
		Signin:   true,
		UserName: userName,
		UserId:   userId,
	}

	json.NewEncoder(w).Encode(res)

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

// ClientSiging ... for new signin
func (d *DataBase) ClientSiging(c *walk.ClientSignin, w http.ResponseWriter) {
	row := d.Conn.QueryRow(`
			SELECT 
				user_id,
				email,
				password
			FROM
				signin
			WHERE
				email = $1
			`, c.Email,
	)

	var client walk.ClientSignin
	_ = row.Scan(&client.UserId, &client.Email, &client.Password)

	if client.Email == "" || client.UserId == 0 {
		log.Println("not found")
	}

	userName := strings.Split(client.Email, "@")[0]

	msg := walk.Message{
		Signin:   true,
		UserName: userName,
		UserId:   client.UserId,
	}

	json.NewEncoder(w).Encode(&msg)
}

// DeleteAndRefound
func (d *DataBase) DeleteAndRefound(purchaseId string) *walk.OrderRefound {
	var orderRefound walk.OrderRefound
	tx, err := d.Conn.Begin()
	if err != nil {
		log.Println(err)
	}
	defer tx.Rollback()

	row := tx.QueryRow(`
			SELECT 
				ci.card_number,
				ci.cv_number,
				co.total 
			FROM 
				client_info  ci
			JOIN 
				client_order_total co 
			ON 
				ci.purchase_id = co.purchase_id 
			WHERE 
				ci.purchase_id = $1;
	`, purchaseId)

	err = row.Scan(
		&orderRefound.CardNumber,
		&orderRefound.CvNumber,
		&orderRefound.Refound,
	)

	if err != nil {
		log.Println(err)
	}

	_, err = tx.Exec(`
			DELETE FROM 
				client_info 
			WHERE 
				purchase_id = $1`,
		purchaseId,
	)

	if err != nil {
		log.Println(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
	}

	return &orderRefound
}
