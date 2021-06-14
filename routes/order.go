package routes

import (
	"net/http"

	"github.com/redmejia/middleware"
	"github.com/redmejia/order"
)

// Order ... make , delete, update order
func Order(base string, middlewares []middleware.Middlewares) {
	http.HandleFunc(base+"orders", middleware.Use(order.HandleOrder, middlewares...))
}
