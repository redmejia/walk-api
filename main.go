package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/redmejia/connection"
	"github.com/redmejia/cors"
	"github.com/redmejia/middleware"
	"github.com/redmejia/routes"
)

func clear() {
	// clear term
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}
func servRunMsg() {
	fmt.Println("Let's GO 🚀 ")
	fmt.Println("Server is running at http://localhost:8080/v1")
}

const base string = "/v1/"

func main() {
	db, err := connection.Dbconn()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	_ = godotenv.Load()
	middlewares := []middleware.Middlewares{
		middleware.Headers,
		middleware.Logger,
		cors.Cors,
	}
	var fs = http.FileServer(http.Dir(os.Getenv("PIC_PATH_DIR")))
	http.Handle(base+"img/", http.StripPrefix(base+"img/", fs))
	routes.Client(base, middlewares)
	routes.ProductCategories(base, middlewares)
	routes.Product(base, middlewares)
	routes.Order(base, middlewares)
	// clear and run server.
	clear()
	servRunMsg()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
