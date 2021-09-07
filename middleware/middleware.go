package middleware

import (
	"log"
	"net/http"
	"os"

	"github.com/redmejia/logs"
)

func Headers(next http.HandlerFunc) http.HandlerFunc {
	headers := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(headers)
}

func Logger(next http.HandlerFunc) http.HandlerFunc {

	var loger logs.Logers
	loger.Info = log.New(os.Stdout, "REQUEST ", log.Ldate|log.Ltime|log.Lmicroseconds)

	logger := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			loger.Info.Printf("%s 🚚  %s ", r.Host, r.Method)
		} else if r.Method == http.MethodPost {
			loger.Info.Printf("%s 🏗️  %s", r.Host, r.Method)
		} else if r.Method == http.MethodDelete {
			loger.Info.Printf("%s ☠️  %s", r.Host, r.Method)
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(logger)
}
