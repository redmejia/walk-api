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
		_, product, err := dbutils.Retrive(db, dbutils.Product{}, `SELECT * FROM products WHERE product_id = $1`, productID)
		_, size, err := dbutils.Retrive(db, dbutils.Sizes{}, `SELECT * FROM sizes WHERE product_id = $1`, productID)
		_, color, err := dbutils.Retrive(db, dbutils.Colors{}, `SELECT * FROM colors WHERE product_id = $1`, productID)
		if err != nil {
			log.Println("ERROR", err)
		}
		s := size.(dbutils.Sizes)
		c := color.(dbutils.Colors)
		productInfo := dbutils.ProductInfo{
			Product: product.(dbutils.Product),
			Size:    []string{s.SizeOne, s.SizeTwo, s.SizeThree, s.SizeFour},
			Colors:  []string{c.ColorOne, c.ColorTwo, c.ColorThree, c.ColorFour},
		}
		json.NewEncoder(w).Encode(productInfo)
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
