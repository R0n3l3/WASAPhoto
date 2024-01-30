package api

import (
	"encoding/json"
	"github.com/R0n3l3/WASAPhoto/service/database"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	var user database.User
	token := getToken(r.Header.Get("Authorization"))
	user.UserId = token

	name := r.URL.Query().Get("username")

	userId, err := rt.db.CreateUser(name)
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
