package routes

import (
	"net/http"

	"github.com/redmejia/clients"
	"github.com/redmejia/middleware"
)

// Client ... route for client register and signin
func Client(base string, middlewares []middleware.Middleware) {
	http.HandleFunc(base+"register", middleware.Use(clients.HandleRegister, middlewares...))
	http.HandleFunc(base+"signin", middleware.Use(clients.HandlerSignin, middlewares...))
}
