package middleware

import (
	"log"
	"net/http"
)

func Headers(next http.Handler) http.Handler {
	headers := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(headers)
}

func Logger(next http.Handler) http.Handler {
	logger := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			log.Printf("%s 🚚  %s R", r.Host, r.Method)
		} else if r.Method == http.MethodPost {
			log.Printf("%s 🏗️  %s C", r.Host, r.Method)
		} else {
			log.Printf("%s 🔥 %s ", r.Host, r.Method)
		}
		// log.Println(r.Host, r.Method)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(logger)
}
