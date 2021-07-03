package handlers

import (
	"encoding/json"
	"net/http"
)

// http://localhost:8080/v1/product?product-id
// HandleProduct ... retrive product by id
func (s *Store) HandleProduct(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		productID := r.URL.Query().Get("product-id")
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
				products p 
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
				p.product_id = $1
	`
		productInfo := s.Product.GetProductById(query, productID)
		json.NewEncoder(w).Encode(productInfo)
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
