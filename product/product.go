package product

import (
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/redmejia/connection"
	"github.com/redmejia/dbutils"
	"github.com/redmejia/request/queries"
)

// http://localhost:8080/v1/product?cat=mens-boots&pro-id=1
func HandleProducts(w http.ResponseWriter, r *http.Request) {
	var product dbutils.Product
	rQuery := r.URL.Query()
	categorie := rQuery["cat"][0]
	proId := rQuery["pro-id"][0]
	db, err := connection.Dbconn()
	if err != nil {
		log.Println("ERRO ", err)
		return
	}
	defer db.Close()
	switch categorie {
	case queries.MensBoots:
		// concatinating proId to the querry fix error
		query := `SELECT * FROM boots_mens WHERE pro_id = $1`
		data, err := dbutils.Retrive(db, product, query, proId)
		if err != nil {
			log.Println("ERRO ", err)
		}
		json.NewEncoder(w).Encode(data)
	case queries.MensSport:
		query := `SELECT * FROM athletic WHERE pro_id = $1`
		product, err := dbutils.Retrive(db, product, query, proId)
		if err != nil {
			log.Println("ERRO ", err)
		}
		json.NewEncoder(w).Encode(product)
	case queries.WomensBoots:
		query := `SELECT * FROM boots_womens WHERE pro_id = $1`
		product, err := dbutils.Retrive(db, product, query, proId)
		if err != nil {
			log.Println("ERRO ", err)
		}
		json.NewEncoder(w).Encode(product)
	case queries.Heels:
		query := `SELECT * FROM heels WHERE pro_id = $1`
		product, err := dbutils.Retrive(db, product, query, proId)
		if err != nil {
			log.Println("ERRO ", err)
		}
		json.NewEncoder(w).Encode(product)
	default:
		http.Error(w, "Oooops Somethig when wrong. :'( ", http.StatusInternalServerError)
		return
	}
}
