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

// http://localhost:8080/v1/categorie?cat=mens-boots
func HandleCategories(w http.ResponseWriter, r *http.Request) {
	rQ := r.URL.Query().Get("cat")
	switch rQ {
	case queries.MensBoots:
		db, err := connection.Dbconn()
		if err != nil {
			log.Println("ERROR  [-]", err)
		}
		defer db.Close()
		query := `SELECT * FROM boots_mens`
		categorie, err := dbutils.Retrive(db, nil, query, nil)
		if err != nil {
			log.Println("ERROR  [-]", err)
		}
		json.NewEncoder(w).Encode(categorie)
	// case queries.MensSport:
	// 	db, err := connection.Dbconn()
	// 	if err != nil {
	// 		log.Println("ERROR  [-]", err)
	// 	}
	// 	defer db.Close()
	// 	query := `SELECT * FROM athletic`
	// 	categorie, err := dbutils.Retrive(db, query)
	// 	if err != nil {
	// 		log.Println("ERROR  [-]", err)
	// 	}
	// 	json.NewEncoder(w).Encode(categorie)
	// case queries.WomensBoots:
	// 	db, err := connection.Dbconn()
	// 	if err != nil {
	// 		log.Println("ERROR  [-]", err)
	// 	}
	// 	defer db.Close()
	// 	query := `SELECT * FROM boots_womens`
	// 	categorie, err := dbutils.Retrive(db, query)
	// 	if err != nil {
	// 		log.Println("ERROR  [-]", err)
	// 	}
	// 	json.NewEncoder(w).Encode(categorie)
	// case queries.Heels:
	// 	db, err := connection.Dbconn()
	// 	if err != nil {
	// 		log.Println("ERROR  [-]", err)
	// 	}
	// 	defer db.Close()
	// 	query := `SELECT * FROM heels`
	// 	categorie, err := dbutils.Retrive(db, query)
	// 	if err != nil {
	// 		log.Println("ERROR  [-]", err)
	// 	}
	// 	json.NewEncoder(w).Encode(categorie)
	default:
		http.Error(w, "SOMETHIG GOES WRONG", http.StatusInternalServerError)
		return
	}
}
