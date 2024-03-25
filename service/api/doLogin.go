package api

import (
	"encoding/json"
	"net/http"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var name string
	err := json.NewDecoder(r.Body).Decode(&name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userId, err := rt.db.DoLogin(name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}
