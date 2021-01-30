package categories

import (
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/redmejia/connection"
)

const (
	// mens
	mensBoots = "mens-boots"
	mensSport = "mens-sport"
	// womens
	womensBoots = "womens-boots"
	heels       = "heels"
)

// http://localhost:8080/v1/categorie?cat=mens-boots
func Categories(w http.ResponseWriter, r *http.Request) {
	rQ := r.URL.Query().Get("cat")
	switch rQ {
	case mensBoots:
		db, err := connection.Dbconn()
		if err != nil {
			log.Println("ERROR  [-]", err)
			return
		}
		defer db.Close()
		query := `SELECT * FROM boots_mens`
		products, err := retriveProducts(db, query)
		if err != nil {
			return
		}
		json.NewEncoder(w).Encode(products)
	case mensSport:
		db, err := connection.Dbconn()
		if err != nil {
			log.Println("ERROR  [-]", err)
			return
		}
		defer db.Close()
		query := `SELECT * FROM athletic`
		products, err := retriveProducts(db, query)
		if err != nil {
			return
		}
		json.NewEncoder(w).Encode(products)
	case womensBoots:
		db, err := connection.Dbconn()
		if err != nil {
			log.Println("ERROR  [-]", err)
			return
		}
		defer db.Close()
		query := `SELECT * FROM boots_womens`
		products, err := retriveProducts(db, query)
		if err != nil {
			return
		}
		json.NewEncoder(w).Encode(products)
	case heels:
		db, err := connection.Dbconn()
		if err != nil {
			log.Println("ERROR  [-]", err)
			return
		}
		defer db.Close()
		query := `SELECT * FROM heels`
		products, err := retriveProducts(db, query)
		if err != nil {
			return
		}
		json.NewEncoder(w).Encode(products)
	default:
		http.Error(w, "SOMETHIG GOES WRONG", http.StatusInternalServerError)
		return
	}
}
