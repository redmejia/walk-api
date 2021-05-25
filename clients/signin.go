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
	switch r.Method {
	case http.MethodPost:
		db, err := connection.Dbconn()
		if err != nil {
			log.Fatal(err)
		}
		var signin dbutils.Signin
		json.NewDecoder(r.Body).Decode(&signin)
		query := `SELECT user_id, email, password from signin WHERE email = $1`
		email := signin.Email
		client, err := dbutils.Retrive(db, signin, query, email)
		if err != nil {
			log.Fatal(err)
		}
		if len(client) == 0 {
			fmt.Println("not found...")
		} else {
			cl := client[0].(dbutils.Signin) // asserting
			if cl.Email == signin.Email && cl.Password == signin.Password {
				var res = struct {
					Signin bool `json:"signin"`
					UserId int  `json:"user_id"` // add more filds as required
				}{
					Signin: true,
					UserId: cl.UserId,
				}

				json.NewEncoder(w).Encode(res)
			}
		}
	case http.MethodOptions:
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
