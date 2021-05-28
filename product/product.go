package product

import (
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/redmejia/connection"
	"github.com/redmejia/dbutils"
)

// NOT NEED ????
// http://localhost:8080/v1/product?cat=mens-boots&pro-id=1
// func HandleProducts(w http.ResponseWriter, r *http.Request) {
// 	var product dbutils.Products
// 	rQuery := r.URL.Query()
// 	categorie := rQuery["cat"][0]
// 	proId := rQuery["pro-id"][0]
// 	db, err := connection.Dbconn()
// 	if err != nil {
// 		log.Println("ERRO ", err)
// 		return
// 	}
// 	defer db.Close()
// 	switch categorie {
// 	case queries.MensBoots:
// 		// concatinating proId to the querry fix error
// 		query := `SELECT * FROM boots_mens WHERE pro_id = $1`
// 		data, err := dbutils.Retrive(db, product, query, proId)
// 		if err != nil {
// 			log.Println("ERRO ", err)
// 		}
// 		json.NewEncoder(w).Encode(data)
// 	case queries.MensSport:
// 		query := `SELECT * FROM athletic WHERE pro_id = $1`
// 		product, err := dbutils.Retrive(db, product, query, proId)
// 		if err != nil {
// 			log.Println("ERRO ", err)
// 		}
// 		json.NewEncoder(w).Encode(product)
// 	case queries.WomensBoots:
// 		query := `SELECT * FROM boots_womens WHERE pro_id = $1`
// 		product, err := dbutils.Retrive(db, product, query, proId)
// 		if err != nil {
// 			log.Println("ERRO ", err)
// 		}
// 		json.NewEncoder(w).Encode(product)
// 	case queries.Heels:
// 		query := `SELECT * FROM heels WHERE pro_id = $1`
// 		product, err := dbutils.Retrive(db, product, query, proId)
// 		if err != nil {
// 			log.Println("ERRO ", err)
// 		}
// 		json.NewEncoder(w).Encode(product)
// 	default:
// 		http.Error(w, "Oooops Somethig when wrong. :'( ", http.StatusInternalServerError)
// 		return
// 	}
// }

// http://localhost:8080/v1/product?product-id
func HandleProducts(w http.ResponseWriter, r *http.Request) {
	db, err := connection.Dbconn()
	if err != nil {
		log.Println("ERRO ", err)
		return
	}
	defer db.Close()
	productID := r.URL.Query().Get("product-id")
	// productID, _ := strconv.Atoi(rQuery["product-d"])
	query := `SELECT * FROM products WHERE product = $1`
	data, err := dbutils.Retrive(db, dbutils.Product{}, query, productID)
	if err != nil {
		log.Println("ERRO ", err)
	}
	json.NewEncoder(w).Encode(data)
}
