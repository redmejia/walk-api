package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/redmejia/request"
)

func clear() {
	// clear term
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Let's Hack ")
}

func main() {
	// serv images
	// this will chage need to work review
	var fs = http.FileServer(http.Dir("/home/red/Desktop/users")) // path on env change
	http.Handle("/v1/img/", http.StripPrefix("/v1/img/", fs))
	http.HandleFunc("/v1", root)
	http.Handle("/v1/register", request.Register)
	http.Handle("/v1/signin", request.Signin)
	http.Handle("/v1/categorie", request.Catergories)
	http.Handle("/v1/orders", request.Order)
	http.Handle("/v1/product", request.Product)
	// http.Handle("/v1/promos", request.Promos)
	http.Handle("/v1/promo", request.Promo) // this will change
	// clear and run server.
	clear()
	fmt.Println("Let's GO 🚀 ")
	fmt.Println("Server is running at http://localhost:8080/v1")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
