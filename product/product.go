package product

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/redmejia/connection"
	"github.com/redmejia/dbutils"
)

const (
	// mens
	mensBoots = "mens-boots"
	mensSport = "mens-sport"
	// womens
	womensBoots = "womens-boots"
	heels       = "heels"
)

// http://localhost:8080/v1/product?cat=mens-boots&pro-id=1
func HandleProducts(w http.ResponseWriter, r *http.Request) {
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
	case mensBoots:
		query := `SELECT * FROM boots_mens WHERE pro_id = ` + proId
		product, err := dbutils.Retrive(db, query)
		if err != nil {
			log.Println("ERRO ", err)
		}
		json.NewEncoder(w).Encode(product)
	case mensSport:
		query := `SELECT * FROM athletic WHERE pro_id = ` + proId
		product, err := dbutils.Retrive(db, query)
		if err != nil {
			log.Println("ERRO ", err)
		}
		json.NewEncoder(w).Encode(product)
	case womensBoots:
		query := `SELECT * FROM boots_womens WHERE pro_id = ` + proId
		product, err := dbutils.Retrive(db, query)
		if err != nil {
			log.Println("ERRO ", err)
		}
		json.NewEncoder(w).Encode(product)
	case heels:
		query := `SELECT * FROM heels WHERE pro_id = ` + proId
		product, err := dbutils.Retrive(db, query)
		if err != nil {
			log.Println("ERRO ", err)
		}
		json.NewEncoder(w).Encode(product)
	default:
		http.Error(w, "Oooops Somethig when wrong. :'( ", http.StatusInternalServerError)
		return
	}
}
