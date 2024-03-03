package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	name := r.URL.Query().Get("username")

	userId, err := rt.db.DoLogin(name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
