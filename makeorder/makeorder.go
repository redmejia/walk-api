package makeorder

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/redmejia/connection"
	"github.com/redmejia/dbutils"
)

func Makeorder(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		db, err := connection.Dbconn()
		if err != nil {
			log.Println("ERROR  [-]", err)
			return
		}
		defer db.Close()
		var order dbutils.Order
		err = json.NewDecoder(r.Body).Decode(&order)
		if err != nil {
			fmt.Println("ERROR ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = dbutils.NewOrder(db, order.ProID, order.Name, order.Color, order.Size, order.Total)
		if err != nil {
			fmt.Println("ERROR INSERT ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("new record was created."))
	}
}
