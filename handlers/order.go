package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	ms_or "github.com/redmejia/order"
	"github.com/redmejia/walk"
)

func (s *Store) HandleOrder(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var order walk.ClientOrder
		err := json.NewDecoder(r.Body).Decode(&order)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("recives order")
		fmt.Println(order)
		client := walk.ClientCardInfo{
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
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		var status walk.PurchaseStatus
		json.NewDecoder(resp.Body).Decode(&status)
		if status.TransactionCode == 0 {
			msg := ms_or.Message{
				PurchaseMSG: "No matching card",
				Suggestion:  "Verify your card number",
			}
			json.NewEncoder(w).Encode(msg)
		} else if status.TransactionCode == 2 || status.TransactionCode == 5 {
			// newOrder(&order, &status)
			order.NewOrder(&status)
		}
	case http.MethodGet:
		// http://localhost:8080/v1/orders?user-id=2
		return
		// uid, _ := strconv.Atoi(r.URL.Query().Get("user-id"))
		// purchase := clientPurchase(uid)
		// json.NewEncoder(w).Encode(purchase)
	case http.MethodOptions:
		return
	}
}
