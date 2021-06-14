package categories

import (
	"encoding/json"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/redmejia/dbutils"
)

// This are the four categories constans for a enter query
// http://localhost:8080/v1/categorie?cat=mens-boots
const (
	// mens
	MensBoots = "mens-boots"
	MensSport = "mens-sport"
	// womens
	WomensBoots = "womens-boots"
	Heels       = "heels"
)

// http://localhost:8080/v1/product?cat=mens-boots
func HandleCategories(w http.ResponseWriter, r *http.Request) {
	rQ := r.URL.Query().Get("cat")
	var product dbutils.Products
	switch rQ {
	case MensBoots:
		product, _ := dbutils.Retrive(product, `
					select
						p.product_id,
						p.pro_name,
						p.price,
						i.img_one_path
					from
						boots_mens p
					join
						shoes_img i
					on
						p.product_id = i.product_id
		 `)
		json.NewEncoder(w).Encode(product)
	case MensSport:
		product, _ := dbutils.Retrive(product, `
					select
						p.product_id,
						p.pro_name,
						p.price,
						i.img_one_path
					from
						athletic p
					join
						shoes_img i
					on
						p.product_id = i.product_id
		 `)
		json.NewEncoder(w).Encode(product)
	case WomensBoots:
		product, _ := dbutils.Retrive(product, `
					select
						p.product_id,
						p.pro_name,
						p.price,
						i.img_one_path
					from
						boots_womens p
					join
						shoes_img i
					on
						p.product_id = i.product_id
		`)
		json.NewEncoder(w).Encode(product)
	case Heels:
		product, _ := dbutils.Retrive(product, `
					select 
						p.product_id, 
						p.pro_name, 
						p.price, 
						i.img_one_path 
					from 
						heels p 
					join 
						shoes_img i 
					on 
						p.product_id = i.product_id
		`)
		json.NewEncoder(w).Encode(product)
	default:
		http.Error(w, "SOMETHIG GOES WRONG", http.StatusInternalServerError)
		return
	}
}
