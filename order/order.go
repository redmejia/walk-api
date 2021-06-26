package order

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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
		for k, v := range order.Items {
			fmt.Println(k, v.ProName)
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
		return
	case http.MethodOptions:
		return

	}
}

func newOrder(order *Order, status *PurchaseStatus) {
	fmt.Println("*status", *status)
	db := connection.DB

	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
	}
	defer tx.Rollback()
	_, err = tx.Exec(
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
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
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
	)
	if err != nil {
		log.Println("insert client info ", err)
	}

	for _, v := range order.Items {
		_, err = tx.Exec(`
			INSERT INTO client_order(
				user_id,
				product_id,
				pro_name,
				color,
				size,
				qty,
				price)
			VALUES ($1, $2, $3, $4, $5, $6, $7)`,
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
	err = tx.Commit()
	if err != nil {
		log.Println(err)
	}
}
