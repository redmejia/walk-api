package routes

import (
	"fmt"
	"net/http"

	"github.com/redmejia/handlers"
	"github.com/redmejia/middleware"
)

// Product ... for retriving product and promotions
func Product(base string, middlewares []middleware.Middleware) {
	// http.HandleFunc(fmt.Sprintf("%sproduct", base), middleware.Use(handlers.HandleProduct, middlewares...))
	http.HandleFunc(fmt.Sprintf("%spromo", base), middleware.Use(handlers.HandlerPromo, middlewares...))
}
