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
	var signin dbutils.SigninForm
	json.NewDecoder(r.Body).Decode(&signin)
	query := `SELECT email, password from signin WHERE email = $1`
	cliente, err := dbutils.Retrive(db, signin, query, signin.Email)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("email : ", signin.Email)
	fmt.Println("Pasword : ", signin.Password)
	if len(cliente) == 0 {
		fmt.Println("not found...")
	} else {
		fmt.Println("dara : ", cliente)
	}
}
