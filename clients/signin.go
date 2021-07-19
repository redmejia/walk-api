package clients

import (
	"encoding/json"
	"net/http"

	"github.com/redmejia/walk"
)

func HandlerSignin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var signin walk.ClientSignin
		json.NewDecoder(r.Body).Decode(&signin)
		// var store walk.Store = &signin
		signin.Client(w)
		// store.Client(w)
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
