package middleware

import "net/http"

type Middlewares func(http.HandlerFunc) http.HandlerFunc

func Use(handler http.HandlerFunc, list ...Middlewares) http.HandlerFunc {
	if len(list) < 1 {
		return handler
	}
	wrapHandler := handler
	for i := len(list) - 1; i >= 0; i-- {
		wrapHandler = list[i](wrapHandler)

	}

	return wrapHandler
}
