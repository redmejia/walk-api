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
		fmt.Println("recives order")
		fmt.Println(order)
		client := ClientInfo{
			CardNumber:     order.Client.CardNumber,
			CvNumber:       order.Client.CvNumber,
			PurchaseAmount: order.Total,
		}
		clientCardInfo, err := json.Marshal(&client)
		if err != nil {
			fmt.Println(err)
		}
		// request to middle serv
		resp, err := http.Post("http://127.0.0.1:8082/order", "application/json", bytes.NewBuffer(clientCardInfo))
		if err != nil {
			fmt.Println("Erro resp ", err)
			return
		}
		defer resp.Body.Close()
		var status PurchaseStatus
		json.NewDecoder(resp.Body).Decode(&status)
		fmt.Println("test status code ")
		fmt.Println(status)
		if status.TransactionCode == 0 {
			msg := Message{
				PurchaseMSG: "No matching card",
				Suggestion:  "Verify your card number",
			}
			json.NewEncoder(w).Encode(msg)
		} else if status.TransactionCode == 2 || status.TransactionCode == 5 {
			fmt.Println("you are trying")
			fmt.Println(order)
			newOrder(&order, &status)
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

func checkPurchaseStat(userI int) bool {
	db := connection.DB
	var statusCode int
	err := db.QueryRow(`
		SELECT 
			purchase_code 
		FROM 
			client_purchase_status
		WHERE 
			user_id = $1
	`, userI).Scan(&statusCode)
	if err != nil {
		log.Println(err)
	}
	return statusCode == 5
}

func clientPurchase(userId int) (purchase Purchase) {
	db := connection.DB
	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
	}
	defer tx.Rollback()
	// var client []Client
	// rows, err := tx.Query(`
	// 	SELECT
	// 		distinct purchase_id,
	// 		user_id,
	// 		first_name,
	// 		last_name,
	// 		email,
	// 		address,
	// 		state,
	// 		zip
	// 	FROM
	// 		client_info
	// 	WHERE
	// 		user_id = $1
	// 	ORDER BY
	// 		purchase_id
	// `, userId)
	// if err != nil {
	// 	log.Println(err)
	// }
	// defer rows.Close()
	// for rows.Next() {
	// 	var cl Client
	// 	rows.Scan(
	// 		&cl.PurchaseID, &cl.UserId, &cl.FirstName, &cl.LastName, &cl.Email,
	// 		&cl.Address, &cl.State, &cl.Zip,
	// 	)
	// 	client = append(client, cl)
	// }
	var product []MyOrder
	rows, err := tx.Query(`
		SELECT distinct ci.purchase_id,
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
			cs.purchase_code as status_code
		FROM client_info ci
			JOIN client_order o ON ci.user_id = ci.user_id
			AND o.purchase_id = ci.purchase_id
			JOIN shoes_img s ON o.product_id = s.product_id
			JOIN client_purchase_status cs ON o.user_id = cs.user_id
			AND cs.purchase_id = o.purchase_id
		WHERE ci.user_id = $1 
		ORDER BY ci.purchase_id
	`, userId)
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		var pt MyOrder
		rows.Scan(
			&pt.Client.PurchaseID, &pt.Client.UserId, &pt.Client.FirstName, &pt.Client.LastName, &pt.Client.Email,
			&pt.Client.Address, &pt.Client.State, &pt.Client.Zip, &pt.Product.PurchaseID, &pt.Product.ProductId,
			&pt.Product.ProName, &pt.Product.Color, &pt.Product.Size, &pt.Product.Qty,
			&pt.Product.Price, &pt.Product.Img, &pt.Product.StatusCode,
		)
		//rows.Scan(
		//	&pt.PurchaseID, &pt.ProductId, &pt.ProName, &pt.Color,
		//	&pt.Size, &pt.Qty, &pt.Price, &pt.Img, &pt.StatusCode,
		//)
		product = append(product, pt)
	}
	var totals []Totals
	rows, err = tx.Query(`
		SELECT 
			distinct purchase_id,
			total
		FROM 
			client_order_total
		WHERE 
			user_id = $1 
		ORDER BY 
			purchase_id
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
		// Client: client,
		Orders: product,
		Totals: totals,
	}
	err = tx.Commit()
	if err != nil {
		log.Println(err)
	}
	return
}
