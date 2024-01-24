package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Create new user
func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	name := r.URL.Query().Get("username") // Get the name of the user

	var user User
	token := getToken(r.Header.Get("Authorization"))
	user.userId = token

	for i := 0; i <= len(Users); i++ { // If the user exists, return the name
		if Users[i].Username == name {
			err := json.NewEncoder(w).Encode(name)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

		}
	}

	// -Return the name
	err := json.NewEncoder(w).Encode(name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
