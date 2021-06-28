package order

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/redmejia/connection"
)

func HandleOrder(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var order Order
		err := json.NewDecoder(r.Body).Decode(&order)
		if err != nil {
			fmt.Println("ERROR ", err)
			// w.WriteHeader(http.StatusInternalServerError)
			return
		}
		client := ClientInfo{
			CardNumber:     order.Client.CardNumber,
			CvNumber:       order.Client.CvNumber,
			PurchaseAmount: order.Total,
		}
		jdata, err := json.Marshal(&client)
		if err != nil {
			fmt.Println(err)
		}
		// request to middle serv
		resp, err := http.Post("http://127.0.0.1:8082/order", "application/json", bytes.NewBuffer(jdata))
		if err != nil {
			fmt.Println("Erro resp ", err)
			return
		}
		defer resp.Body.Close()
		var status PurchaseStatus
		json.NewDecoder(resp.Body).Decode(&status)
		if status.TransactionCode == 2 {
			newOrder(&order, &status)
			// msg := Message{
			// 	PurchaseMSG: "No matching card",
			// 	Suggestion:  "Verify your card number",
			// }
			// json.NewEncoder(w).Encode(msg)
		}
		// json.NewEncoder(w).Encode(status)
	// fmt.Println("youuu ", msg)
	// _, err = dbutils.NewOrder(order.ProID, order.Name, order.Color, order.Size, order.Total)
	// if err != nil {
	// 	fmt.Println("ERROR INSERT ", err)
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	// w.WriteHeader(http.StatusCreated)
	// w.Write([]byte("new record was created."))
	case http.MethodGet:
		// http://localhost:8080/v1/orders?user-id=2
		uid, _ := strconv.Atoi(r.URL.Query().Get("user-id"))
		purchase := clientPurchase(uid)
		json.NewEncoder(w).Encode(purchase)
	case http.MethodOptions:
		return

	}
}

func newOrder(order *Order, status *PurchaseStatus) {
	db := connection.DB

	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
	}
	defer tx.Rollback()
	var purchaseID int
	err = tx.QueryRow(
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
		order.Client.UserId,
		order.Client.FirstName,
		order.Client.LastName,
		order.Client.Email,
		order.Client.Address,
		order.Client.State,
		order.Client.Zip,
		order.Client.NameOnCard,
		order.Client.CardNumber,
		order.Client.CvNumber,
	).Scan(&purchaseID)
	if err != nil {
		log.Println("insert client info ", err)
	}

	for _, v := range order.Items {
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
			order.Client.UserId,
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
		VALUES ($1, $2, $3)`, purchaseID, order.Client.UserId, order.Total)
	if err != nil {
		log.Println(err)
	}
	_, err = tx.Exec(`INSERT INTO  client_purchase_status(purchase_id, user_id, purchase_status, purchase_code)
			VALUES ($1, $2, $3, $4)`, purchaseID, order.Client.UserId, status.Status, status.TransactionCode)
	if err != nil {
		log.Println(err)
	}
	err = tx.Commit()
	if err != nil {
		log.Println(err)
	}
}

func clientPurchase(userId int) (purchase Purchase) {
	db := connection.DB
	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
	}
	defer tx.Rollback()
	var client []Client
	rows, err := tx.Query(`
		SELECT 
			distinct purchase_id,
			user_id,
			first_name,
			last_name,
			email,
			address,
			state,
			zip
		FROM 
			client_info
		WHERE 
			user_id = $1
		ORDER BY 
			purchase_id
	`, userId)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var cl Client
		rows.Scan(
			&cl.PurchaseID, &cl.UserId, &cl.FirstName, &cl.LastName, &cl.Email,
			&cl.Address, &cl.State, &cl.Zip,
		)
		client = append(client, cl)
	}
	var product []Product
	rows, err = tx.Query(`
		SELECT 
			distinct o.purchase_id as purchase_id,
			o.product_id,
			o.pro_name,
			o.color,
			o.size,
			o.qty,
			o.price,
			s.img_one_path
		FROM 
			client_order o
		JOIN 
			shoes_img s 
		ON 
			o.product_id = s.product_id
		WHERE 
			o.user_id = $1
		ORDER BY 
			purchase_id
	`, userId)
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		var pt Product
		rows.Scan(
			&pt.PurchaseID, &pt.ProductId, &pt.ProName, &pt.Color,
			&pt.Size, &pt.Qty, &pt.Price, &pt.Img,
		)
		product = append(product, pt)
	}
	var totals []Totals
	rows, err = tx.Query(`
		SELECT distinct purchase_id,
			total
		FROM client_order_total
		WHERE user_id = $1 
		ORDER BY purchase_id
	`, userId)
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		var t Totals
		rows.Scan(
			&t.PurchaseID,
			&t.Total,
		)
		totals = append(totals, t)
	}
	purchase = Purchase{
		Client: client,
		Orders: product,
		Totals: totals,
	}
	err = tx.Commit()
	if err != nil {
		log.Println(err)
	}
	return
}
