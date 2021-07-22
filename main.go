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

func main() {
	db, err := connection.Dbconn()

	if err != nil {
		log.Println(err)
	}

	defer db.Close()

	_ = godotenv.Load()

	routes.Routes()

	// clear and run server.
	clear()
	servRunMsg()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
