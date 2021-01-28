package makeorder

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type newOrder struct {
	Name  string  `json:"name"`
	Unit  uint8   `json:"unit"`
	Price float32 `json:"price"`
}

func Makeorder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == http.MethodPost {
		var order newOrder
		err := json.NewDecoder(r.Body).Decode(&order)
		if err != nil {
			fmt.Println("ERROR", err)
			return
		}
		// this must insert new record in db
		fmt.Println(order)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("new record was created."))
	}
}
