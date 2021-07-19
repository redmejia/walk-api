package clients

import (
	"encoding/json"
	"net/http"

	"github.com/redmejia/walk"
)

// HandleRegiter ... register clients
func HandleRegister(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var register walk.ClientRegister
		json.NewDecoder(r.Body).Decode(&register)
		// var store walk.Store = &register
		register.Client(w)
		// store.Client(w)
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
