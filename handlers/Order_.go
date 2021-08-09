package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
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
func (s *StoreHandler) HandleOrder(w http.ResponseWriter, r *http.Request) {
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

		// request POST to small bank restapi
		resp, err := http.Post("http://127.0.0.1:8081/v1/purchase/transaction",
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
			s.Store.InsertNewOrder(&order, status)

		}

	case http.MethodGet:
		// http://localhost:8080/v1/orders?user-id=2
		uid, _ := strconv.Atoi(r.URL.Query().Get("user-id"))
		purchase := s.Store.GetClientPurchaseInfoByUserId(uid)
		json.NewEncoder(w).Encode(purchase)
		return
	case http.MethodDelete:
		// http://localhost:8080/v1/orders?del-refound
		purchaseId := r.URL.Query().Get("del-refound")
		refound := s.Store.DeleteAndRefound(purchaseId)
		log.Println(refound)
	case http.MethodOptions:
		return
	}
}
