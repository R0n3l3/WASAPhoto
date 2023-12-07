package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	newName := r.URL.Query().Get("username") //Get the new username
	oldName := ps.ByName("userId")           //Get the old username

	for i := 0; i <= len(Users); i++ { //Change the username
		if Users[i].Username == oldName {
			Users[i].Username = newName
			Users[i].UserProfile.ProfileId = newName
			err := json.NewEncoder(w).Encode(Users[i])
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}
	}
	w.WriteHeader(http.StatusNotFound) //The user does not exist
	return
}
