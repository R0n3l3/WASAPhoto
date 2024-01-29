package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	myName := ps.ByName("username")
	name := r.URL.Query().Get("username") //Get the user I want to search

	isBanned := rt.db.IsBanned(name, myName)
	if isBanned {
		err := json.NewEncoder(w).Encode("You are banned")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	profile, _ := rt.db.GetUserProfile(name)
	err := json.NewEncoder(w).Encode(profile)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
