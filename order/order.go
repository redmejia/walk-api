package order

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/redmejia/dbutils"
)

type PurchaseStatus struct {
	Status          string `json:"status"`
	TransactionCode uint8  `json:"transaction_code"` // 00 error card num or cv not valid, 02 ok,  05 not enough to compleate 0.0 balance or purchase is grather than amount
}

func HandleOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var order dbutils.Order
		err := json.NewDecoder(r.Body).Decode(&order)
		if err != nil {
			fmt.Println("ERROR ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Println("recived ", order)
		var nOrder = struct {
			CardNumber     string  `json:"card_number"`
			CvNumber       uint8   `json:"cv_number"`
			PurchaseAmount float32 `json:"purchase_amount"`
		}{
			CardNumber:     order.CardNumber,
			CvNumber:       order.CvNumber,
			PurchaseAmount: order.Total,
		}
		jdata, err := json.Marshal(&nOrder)
		if err != nil {
			fmt.Println("err ", err)
		}
		// request to middle serv
		resp, err := http.Post("http://127.0.0.1:8082/order", "application/json", bytes.NewBuffer(jdata))
		if err != nil {
			fmt.Println("Erro resp ", err)
			return
		}
		defer resp.Body.Close()
		fmt.Println("handle order ")
		var status PurchaseStatus
		json.NewDecoder(resp.Body).Decode(&status)
		fmt.Println("handleorder status, ", status.Status)
		// fmt.Println("youuu ", msg)
		// _, err = dbutils.NewOrder(order.ProID, order.Name, order.Color, order.Size, order.Total)
		// if err != nil {
		// 	fmt.Println("ERROR INSERT ", err)
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	return
		// }
		// w.WriteHeader(http.StatusCreated)
		// w.Write([]byte("new record was created."))
	}
}
