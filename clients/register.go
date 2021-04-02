package clients

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/redmejia/connection"
)

type form registerForm

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	db, err := connection.Dbconn()
	if err != nil {
		fmt.Println("ERR ", err)
	}
	defer db.Close()
	var register form
	json.NewDecoder(r.Body).Decode(&register)

	tx, err := db.Begin()

	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	query := `INSERT INTO register (name, email) VALUES ($1, $2)`
	registerStm, err := tx.Prepare(query)
	if err != nil {
		log.Fatal("ERROR on register ", err)
	}
	_, err = registerStm.Exec(register.Name, register.Email)
	if err != nil {
		log.Fatal("EXEC 1 ", err)
	}

	querySig := `INSERT INTO signin (email, password) VALUES ($1, $2)`
	signinStm, err := tx.Prepare(querySig)
	if err != nil {
		log.Fatal("ERROO SIGNIN ", err)
	}
	_, err = signinStm.Exec(register.Email, register.Password)
	if err != nil {
		log.Fatal("EXEC 2 ", err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal("ERRR COMMIT", err)
	}

}
