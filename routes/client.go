package routes

import (
	"fmt"
	"net/http"

	"github.com/redmejia/clients"
	"github.com/redmejia/middleware"
)

// Client ... route for client register and signin
func Client(base string, middlewares []middleware.Middleware) {
	http.HandleFunc(fmt.Sprintf("%sregister", base), middleware.Use(clients.HandleRegister, middlewares...))
	http.HandleFunc(fmt.Sprintf("%ssignin", base), middleware.Use(clients.HandlerSignin, middlewares...))
}
