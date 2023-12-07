package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	name := r.URL.Query().Get("username") //Get the user I want to search

	for i := 0; i <= len(Users); i++ { //When I find the user, return their profile
		if Users[i].Username == name {
			err := json.NewEncoder(w).Encode(Users[i].UserProfile)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}
	}
	w.WriteHeader(http.StatusNotFound) //If I couldn't find the user, error
	return
}
