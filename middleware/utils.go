package middleware

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc

// Middleware
type Middlewares struct {
	Middle []Middleware
}

// SetMiddleware ... set middlewares and cors
func SetMiddleware(middlewares ...Middleware) *Middlewares {
	return &Middlewares{
		Middle: middlewares,
	}
}

// Use middleware wrap chanin middleware
func Use(handler http.HandlerFunc, list ...Middleware) http.HandlerFunc {
	if len(list) < 1 {
		return handler
	}
	wrapHandler := handler
	for i := len(list) - 1; i >= 0; i-- {
		wrapHandler = list[i](wrapHandler)

	}

	return wrapHandler
}
