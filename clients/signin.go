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
		s, _ := dbutils.Retrive(db, dbutils.Signin{}, `
			select 
				user_id, 
				email, 
				password 
			from 
				signin 
			where 
				email = $1`, signin.Email)
		if len(s) == 0 {
			fmt.Println("user not found")
		} else {
			client := s[0].(dbutils.Signin)
			msg := Message{
				Signin: true,
				UserId: client.UserId,
			}
			json.NewEncoder(w).Encode(msg)
		}
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
