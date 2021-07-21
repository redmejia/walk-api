package routes

import "github.com/redmejia/middleware"

// Client ... route for client register and signin
func Client(base string, middlewares []middleware.Middleware) {
	return
	// http.HandleFunc(fmt.Sprintf("%sregister", base), middleware.Use(clients.HandleRegister, middlewares...))
	// http.HandleFunc(fmt.Sprintf("%ssignin", base), middleware.Use(clients.HandlerSignin, middlewares...))
}
