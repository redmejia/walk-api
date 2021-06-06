package categories

import (
	"encoding/json"
	"fmt"
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
	var product dbutils.Products
	switch rQ {
	case queries.MensBoots:
		db, err := connection.Dbconn()
		if err != nil {
			log.Println("ERROR  [-]", err)
		}
		defer db.Close()
		product, _, err := dbutils.Retrive(db, product, `SELECT * FROM boots_mens`)
		if err != nil {
			log.Println("ERROR  [-]", err)
		}
		fmt.Println("products ", product)
		json.NewEncoder(w).Encode(product)
	case queries.MensSport:
		db, err := connection.Dbconn()
		if err != nil {
			log.Println("ERROR  [-]", err)
		}
		defer db.Close()
		product, _, err := dbutils.Retrive(db, product, `SELECT * FROM athletic`)
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
		// product, _, err := dbutils.Retrive(db, product, `SELECT * FROM boots_womens`)
		product, _, err := dbutils.Retrive(db, product, `select p.product_id, p.pro_name, p.price, i.img_one_path from boots_womens p join shoes_img i on p.product_id = i.product_id`)
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
		product, _, err := dbutils.Retrive(db, product, `SELECT * FROM heels`)
		if err != nil {
			log.Println("ERROR  [-]", err)
		}
		json.NewEncoder(w).Encode(product)
	default:
		http.Error(w, "SOMETHIG GOES WRONG", http.StatusInternalServerError)
		return
	}
}
