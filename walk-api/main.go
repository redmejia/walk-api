package main

import (
	"fmt"
	"log"
	"net/http"
)

func Categories(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func main() {

	http.HandleFunc("/", Categories)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
