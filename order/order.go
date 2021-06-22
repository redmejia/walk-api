package order

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/redmejia/dbutils"
)

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
		client := ClientInfo{
			CardNumber:     order.CardNumber,
			CvNumber:       order.CvNumber,
			PurchaseAmount: order.Total,
		}
		jdata, err := json.Marshal(&client)
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
