package product

import (
	"fmt"
	"net/http"
)

func ProductHandle(w http.ResponseWriter, r *http.Request) {
	// rQuery := r.URL.Query().Get("pro")
	// db, err := connection.Dbconn()
	// if err != nil {
	// 	log.Println("ERRO ", err)
	// 	return
	// }
	name := "boots"
	query := `SELECT * FROM boots_mens` + name + ` AND last`
	fmt.Println(query)
}
