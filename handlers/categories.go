package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	// mens
	MensBoots = "mens-boots"
	MensSport = "mens-sport"
	// womens
	WomensBoots = "womens-boots"
	Heels       = "heels"
)

// http://localhost:8080/v1/product?cat=mens-boots
// HandleCategories ... Retrive categories
func (s *Store) HandleCategories(w http.ResponseWriter, r *http.Request) {
	rQ := r.URL.Query().Get("cat")
	// var product dbutils.Products
	switch rQ {
	case MensBoots:
		product, err := s.Categorie.GetProducts(`
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
		if err != nil {
			log.Println(err)
			return
		}
		json.NewEncoder(w).Encode(product)
	case MensSport:
		product, err := s.Categorie.GetProducts(`
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
		product, err := s.Categorie.GetProducts(`
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
		product, err := s.Categorie.GetProducts(`
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
