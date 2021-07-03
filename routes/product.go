package routes

import (
	"net/http"

	"github.com/redmejia/handlers"
	"github.com/redmejia/middleware"
)

// Product ... for retriving product and promotions
func Product(base string, middlewares []middleware.Middlewares) {
	var walk handlers.Store
	http.HandleFunc(base+"product", middleware.Use(walk.HandleProduct, middlewares...))
	http.HandleFunc(base+"promo", middleware.Use(walk.HandlerPromo, middlewares...))
}
