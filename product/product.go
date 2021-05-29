package product

import (
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/redmejia/connection"
	"github.com/redmejia/dbutils"
)

// http://localhost:8080/v1/product?product-id
// HandleProducts ...
func HandleProducts(w http.ResponseWriter, r *http.Request) {
	db, err := connection.Dbconn()
	if err != nil {
		log.Println("ERRO ", err)
		return
	}
	defer db.Close()
	switch r.Method {
	case http.MethodGet:
		productID := r.URL.Query().Get("product-id")
		query := `SELECT * FROM products WHERE product = $1`
		data, err := dbutils.Retrive(db, dbutils.Product{}, query, productID)
		if err != nil {
			log.Println("ERRO ", err)
		}
		json.NewEncoder(w).Encode(data)
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
