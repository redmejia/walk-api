package order

import (
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
		_, err = dbutils.NewOrder(order.ProID, order.Name, order.Color, order.Size, order.Total)
		if err != nil {
			fmt.Println("ERROR INSERT ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("new record was created."))
	}
}
