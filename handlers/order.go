package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/redmejia/walk"
)

// PurchaseFailMsg ... Fail card was not found
type PurchaseFailMsg struct {
	PurchaseMSG string `json:"purchase_msg"`
	Suggestion  string `json:"suggestion"`
}

// HandleOrder ... Handle new order and retrive order
func HandleOrder(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var order walk.ClientOrder
		err := json.NewDecoder(r.Body).Decode(&order)

		if err != nil {
			fmt.Println(err)
			return
		}

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
		resp, err := http.Post("http://127.0.0.1:8082/order",
			"application/json", bytes.NewBuffer(clientCardInfo),
		)
		if err != nil {
			fmt.Println(err)
			return
		}

		defer resp.Body.Close()

		var status walk.PurchaseStatus
		json.NewDecoder(resp.Body).Decode(&status)

		if status.TransactionCode == 0 {

			msg := PurchaseFailMsg{
				PurchaseMSG: "No matching card",
				Suggestion:  "Verify your card number",
			}

			json.NewEncoder(w).Encode(msg)

		} else if status.TransactionCode == 2 || status.TransactionCode == 5 {
			// Insert new client order with status of 2 = Aproved or 5 = decline
			// var store walk.Store = &order
			// store.InsertNewOrder(status)
			order.InsertNewOrder(status)
		}

	case http.MethodGet:
		// http://localhost:8080/v1/orders?user-id=2
		uid, _ := strconv.Atoi(r.URL.Query().Get("user-id"))
		var order walk.Order
		purchase := order.GetClientPurchaseInfoByUserId(uid)
		json.NewEncoder(w).Encode(purchase)

	case http.MethodOptions:
		return
	}
}
