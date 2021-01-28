package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/redmejia/categories"
	"github.com/redmejia/makeorder"
)

func clear() {
	// clear term
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}
func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Let's Hack")
}

func main() {
	http.HandleFunc("/v1", root)
	http.HandleFunc("/v1/categorie", categories.Categories)
	http.HandleFunc("/v1/new-order", makeorder.Makeorder)
	// clear and run server.
	clear()
	fmt.Println("Let's GO")
	fmt.Println("Server is running at http://localhost:8080/v1")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
