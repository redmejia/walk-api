package makeorder

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/redmejia/connection"
)

func Makeorder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == http.MethodPost {
		db, err := connection.Dbconn()
		if err != nil {
			log.Println("ERROR  [-]", err)
			return
		}
		defer db.Close()
		var order Product
		err = json.NewDecoder(r.Body).Decode(&order)
		if err != nil {
			fmt.Println("ERROR ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = newOrder(db, order.ProID, order.Name, order.Color, order.Size, order.Total)
		if err != nil {
			fmt.Println("ERROR INSERT ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("new record was created."))
	}
}
