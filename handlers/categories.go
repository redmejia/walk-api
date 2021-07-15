package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/redmejia/walk"
)

const (
	// mens
	MensBoots = "mens-boots"
	MensSport = "mens-sport"
	// womens
	WomensBoots = "womens-boots"
	Heels       = "heels"
)

type HandlerRep struct {
	DBRep DBRepo
}

// http://localhost:8080/v1/categorie?cat=mens-boots
// HandleCategories ... Retrive categories
func (h *HandlerRep) HandleCategories(w http.ResponseWriter, r *http.Request) {
	rQ := r.URL.Query().Get("cat")
	var products walk.Products
	switch rQ {
	case MensBoots:
		pr, err := h.DBRep.GetProducts(`
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
		// product, err := products.GetProducts(`
		// 			select
		// 				p.product_id,
		// 				p.pro_name,
		// 				p.price,
		// 				i.img_one_path
		// 			from
		// 				boots_mens p
		// 			join
		// 				shoes_img i
		// 			on
		// 				p.product_id = i.product_id
		// `)
		if err != nil {
			log.Println(err)
			return
		}
		json.NewEncoder(w).Encode(pr)
	case MensSport:
		product, err := products.GetProducts(`
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
		if err != nil {
			log.Println(err)
		}
		json.NewEncoder(w).Encode(product)
	case WomensBoots:
		product, err := products.GetProducts(`
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
		if err != nil {
			log.Println(err)
		}
		json.NewEncoder(w).Encode(product)
	case Heels:
		product, err := products.GetProducts(`
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
		if err != nil {
			log.Println(err)
		}
		json.NewEncoder(w).Encode(product)
	default:
		http.Error(w, "SOMETHIG GOES WRONG", http.StatusInternalServerError)
		return
	}
}
