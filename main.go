package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/redmejia/categories"
)

func main() {
	http.HandleFunc("/v1", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Let's Hack"))
	})
	http.HandleFunc("/v1/categorie", categories.Categories)
	fmt.Println("Let's GO")
	fmt.Println("Server is running at http://localhost:8080/v1")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
