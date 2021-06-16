package order

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/redmejia/dbutils"
)

type Msg struct {
	Msq    string `json:"msg"`
	Recive bool   `json:"recive"`
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
			FirstName  string `json:"first_name"`
			LastName   string `json:"last_name"`
			CardNumber string `json:"card_number"`
			CvNumber   uint8  `json:"cv_number"`
		}{
			FirstName:  order.FirstName,
			LastName:   order.LastName,
			CardNumber: order.CardNumber,
			CvNumber:   order.CvNumber,
		}
		jdata, err := json.Marshal(&nOrder)
		if err != nil {
			fmt.Println("err ", err)
		}
		resp, err := http.Post("http://192.168.1.104:8082/order", "application/json", bytes.NewBuffer(jdata))
		if err != nil {
			fmt.Println("Erro resp ", err)
		}
		defer resp.Body.Close()
		var msg Msg
		json.NewDecoder(resp.Body).Decode(&msg)
		fmt.Println("youuu ", msg)
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
