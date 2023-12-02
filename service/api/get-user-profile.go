package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	name := r.URL.Query().Get("username")

	for i := 0; i <= len(Users); i++ {
		if Users[i].Username == name {
			json.NewEncoder(w).Encode(Profiles[i])
			return
		}
	}
}
