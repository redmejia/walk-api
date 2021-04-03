package categories

import (
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/redmejia/connection"
	"github.com/redmejia/dbutils"
	"github.com/redmejia/request/queries"
)

// http://localhost:8080/v1/product?cat=mens-boots
func HandleCategories(w http.ResponseWriter, r *http.Request) {
	rQ := r.URL.Query().Get("cat")
	var product dbutils.Product
	switch rQ {
	case queries.MensBoots:
		db, err := connection.Dbconn()
		if err != nil {
			log.Println("ERROR  [-]", err)
		}
		defer db.Close()
		product, err := dbutils.Retrive(db, product, `SELECT * FROM boots_mens`)
		if err != nil {
			log.Println("ERROR  [-]", err)
		}
		json.NewEncoder(w).Encode(product)
	case queries.MensSport:
		db, err := connection.Dbconn()
		if err != nil {
			log.Println("ERROR  [-]", err)
		}
		defer db.Close()
		product, err := dbutils.Retrive(db, product, `SELECT * FROM athletic`)
		if err != nil {
			log.Println("ERROR  [-]", err)
		}
		json.NewEncoder(w).Encode(product)
	case queries.WomensBoots:
		db, err := connection.Dbconn()
		if err != nil {
			log.Println("ERROR  [-]", err)
		}
		defer db.Close()
		product, err := dbutils.Retrive(db, product, `SELECT * FROM boots_womens`)
		if err != nil {
			log.Println("ERROR  [-]", err)
		}
		json.NewEncoder(w).Encode(product)
	case queries.Heels:
		db, err := connection.Dbconn()
		if err != nil {
			log.Println("ERROR  [-]", err)
		}
		defer db.Close()
		product, err := dbutils.Retrive(db, product, `SELECT * FROM heels`)
		if err != nil {
			log.Println("ERROR  [-]", err)
		}
		json.NewEncoder(w).Encode(product)
	default:
		http.Error(w, "SOMETHIG GOES WRONG", http.StatusInternalServerError)
		return
	}
}
