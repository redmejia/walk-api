package promotion

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/redmejia/connection"
	"github.com/redmejia/dbutils"
)

// HandlerPromo ...
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
	db, err := connection.Dbconn()
	if err != nil {
		log.Println("ERRO ", err)
		return
	}
	defer db.Close()
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
		promo, err := dbutils.Retrive(db, dbutils.Products{}, query)
		if err != nil {
			log.Println(err)
			return
		}
		json.NewEncoder(w).Encode(promo)
	} else if productID, ok := rquery["product-id"]; ok {
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
		productInfo := dbutils.RetriveById(db, query, productID[0])
		json.NewEncoder(w).Encode(productInfo)
	}
}
