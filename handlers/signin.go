package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/redmejia/walk"
)

func (s *StoreHandlers) HandlerSignin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var signin walk.ClientSignin
		json.NewDecoder(r.Body).Decode(&signin)
		s.Store.ClientSiging(&signin, w)

	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
