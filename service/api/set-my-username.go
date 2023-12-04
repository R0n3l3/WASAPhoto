package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	newName := r.URL.Query().Get("username")
	oldName := ps.ByName("userId")

	for i := 0; i <= len(Users); i++ {
		if Users[i].Username == oldName {
			Users[i].Username = newName
			Users[i].UserProfile.ProfileId = newName
			json.NewEncoder(w).Encode(Users[i])
		}
	}

}
