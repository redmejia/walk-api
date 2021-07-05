package routes

import (
	"fmt"
	"net/http"

	"github.com/redmejia/handlers"
	"github.com/redmejia/middleware"
)

// Order ... make , delete, update order
func Order(base string, middlewares []middleware.Middlewares) {
	http.HandleFunc(fmt.Sprintf("%sorders", base), middleware.Use(handlers.HandleOrder, middlewares...))
}
