package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/redmejia/walk"
)

func (s *StoreHandlers) HandleRegister(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var register walk.ClientRegister
		json.NewDecoder(r.Body).Decode(&register)
		s.Store.ClientRegister(&register, w)

	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
