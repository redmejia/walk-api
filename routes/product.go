package routes

import (
	"net/http"

	"github.com/redmejia/middleware"
	"github.com/redmejia/product"
	"github.com/redmejia/promotion"
)

// Product ... for retriving product and promotions
func Product(base string, middlewares []middleware.Middlewares) {
	http.HandleFunc(base+"product", middleware.Use(product.HandleProduct, middlewares...))
	http.Handle("/v1/promo", middleware.Use(promotion.HandlerPromo, middlewares...))
}
