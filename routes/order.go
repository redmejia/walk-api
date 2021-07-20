package routes

import (
	"github.com/redmejia/middleware"
)

// Order ... make , delete, update order
func Order(base string, middlewares []middleware.Middleware) {
	// http.HandleFunc(fmt.Sprintf("%sorders", base), middleware.Use(handlers.HandleOrder, middlewares...))
}
