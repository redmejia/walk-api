package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func (s *StoreHandlers) HandlerPromo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleRouteQuery(s, w, r)
	case http.MethodOptions:
		return
	}
}

// Only for retrive one product from promotion
// http://localhost:8080/v1/promo?product-id=2341

// This two ways for retiving products slice
// http://localhost:8080/v1/promo?products
// http://localhost:8080/v1/promo
func handleRouteQuery(s *StoreHandlers, w http.ResponseWriter, r *http.Request) {
	rquery := r.URL.Query()

	// Check if request query map has product-id then retrive product promotion by product id
	if productID, ok := rquery["product-id"]; ok {
		query := `
	 		select
	 			p.product_id,
	 			p.pro_name,
	 			p.price,
	 			s.size_one,
	 			s.size_two,
	 			s.size_three,
	 			s.size_four,
	 			c.color_one,
	 			c.color_two,
	 			c.color_three,
	 			c.color_four,
	 			i.img_one_path,
	 			i.img_two_path
	 		from
	 			promos p
	 		join
	 			sizes s
	 		on
	 			p.product_id = s.product_id
	 		join
	 			colors c
	 		on
	 			c.product_id = p.product_id
	 		join
	 			shoes_img i
	 		on
	 			p.product_id = i.product_id
	 		where
	 			p.product_id = $1`

		productId, _ := strconv.Atoi(productID[0])
		productInfo := s.Store.GetProductById(query, productId)
		json.NewEncoder(w).Encode(productInfo)

	} else {
		query := `
	 	 	select
	 	 		p.product_id,
	 	 		p.pro_name,
	 	 		p.price,
	 	 		i.img_one_path
	 	 	from
	 	 		promos p
	 	 	join
	 	 		shoes_img i
	 	 	on
	 	 		p.product_id = i.product_id`

		product, err := s.Store.GetProducts(query)

		if err != nil {
			log.Println(err)
			return
		}

		json.NewEncoder(w).Encode(product)
	}

}
