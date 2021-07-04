package routes

import (
	"net/http"

	"github.com/redmejia/handlers"
	"github.com/redmejia/middleware"
)

// Order ... make , delete, update order
func Order(base string, middlewares []middleware.Middlewares) {
	var walk handlers.Store
	http.HandleFunc(base+"orders", middleware.Use(walk.HandleOrder, middlewares...))
}
