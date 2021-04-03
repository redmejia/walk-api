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
	registerStm, err := tx.Prepare(`INSERT INTO register (name, email) VALUES ($1, $2)`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = registerStm.Exec(register.Name, register.Email)
	if err != nil {
		log.Fatal(err)
	}
	signinStm, err := tx.Prepare(`INSERT INTO signin (email, password) VALUES ($1, $2)`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = signinStm.Exec(register.Email, register.Password)
	if err != nil {
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
