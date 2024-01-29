package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) setMyUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	newName := r.URL.Query().Get("username") // Get the new username
	oldName := ps.ByName("userId")           // Get the old username

	user, _ := rt.db.SetMyUsername(oldName, newName)
	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return

	}
	return
}
