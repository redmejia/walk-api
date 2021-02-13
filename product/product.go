package product

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/redmejia/connection"
)

func HandleProduct(w http.ResponseWriter, r *http.Request) {
	rQuery := r.URL.Query().Get("pro-id")
	db, err := connection.Dbconn()
	if err != nil {
		log.Println("ERRO ", err)
		return
	}
	defer db.Close()
	query := `SELECT * FROM boots_mens where pro_id = ` + rQuery

	product, err := retriveProduct(db, query)
	if err != nil {
		log.Println("ERRO ", err)
	}

	json.NewEncoder(w).Encode(product)
}
