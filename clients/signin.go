package clients

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/redmejia/connection"
	"github.com/redmejia/dbutils"
)

func HandlerSignin(w http.ResponseWriter, r *http.Request) {
	db, err := connection.Dbconn()
	if err != nil {
		log.Fatal(err)
	}
	var signin dbutils.Signin
	json.NewDecoder(r.Body).Decode(&signin)
	query := `SELECT email, password from signin WHERE email = $1`
	email := signin.Email
	client, err := dbutils.Retrive(db, signin, query, email)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("email : ", signin.Email)
	fmt.Println("Pasword : ", signin.Password)
	if len(client) == 0 {
		fmt.Println("not found...")
	} else {
		fmt.Println("dara : ", client)
	}
}
