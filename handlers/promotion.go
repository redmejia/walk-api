package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/redmejia/walk"
)

func HandlerPromo(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleRouteQuery(w, r)
	case http.MethodOptions:
		return
	}
}

// Only for retrive one product
// http://localhost:8080/v1/promo?product-id=2341
// This two ways for retiving products slice
// http://localhost:8080/v1/promo?products=true
// http://localhost:8080/v1/promo?products
func handleRouteQuery(w http.ResponseWriter, r *http.Request) {
	rquery := r.URL.Query()

	var products walk.Products
	// Check if request query map has key products then retrive all products in promotion
	if _, ok := rquery["products"]; ok {
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

		promo, err := products.GetProducts(query)

		if err != nil {
			log.Println(err)
			return
		}

		json.NewEncoder(w).Encode(promo)
	}

	var product walk.ProductInfo
	// Check if request query map has product-id then retrive product promotion
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

		productInfo := product.GetProductById(query, productID[0])

		json.NewEncoder(w).Encode(productInfo)
	}

}
