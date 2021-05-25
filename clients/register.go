package clients

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/redmejia/connection"
	"github.com/redmejia/dbutils"
)

type register dbutils.Register

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		db, err := connection.Dbconn()
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		var register register
		json.NewDecoder(r.Body).Decode(&register)
		tx, err := db.Begin()

		if err != nil {
			log.Fatal(err)
		}
		defer tx.Rollback()

		var userId int

		err = tx.QueryRow(`
				INSERT INTO 
					register (name, email) 
				VALUES ($1, $2) 
				RETURNING user_id`, register.Name, register.Email).Scan(&userId)
		if err != nil {
			log.Fatal(err)
		}
		// return email, and user id maybe
		signinStm, err := tx.Prepare(`
				INSERT INTO 
					signin (user_id, email, password) 
				VALUES ($1, $2, $3)`)
		if err != nil {
			log.Fatal(err)
		}
		_, err = signinStm.Exec(userId, register.Email, register.Password)
		if err != nil {
			log.Fatal(err)
		}

		res := Message{
			Signin: true,
			UserId: userId,
		}
		json.NewEncoder(w).Encode(res)
		err = tx.Commit()
		if err != nil {
			log.Fatal(err)
		}
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
