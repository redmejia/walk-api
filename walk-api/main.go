package main

import (
	"log"
	"net/http"

	"github.com/redmejia/categories"
)

func main() {
	http.HandleFunc("/categorie", categories.Categories)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
