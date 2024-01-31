package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	myName := ps.ByName("userId")          //Get my name
	toBan := r.URL.Query().Get("username") //Get the name of the banned person

	isAuth := rt.db.IsAuthorized(getToken(r.Header.Get("Authorization")))
	if !isAuth {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	user, _ := rt.db.BanUser(toBan, myName)
	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return

	}
	err = rt.db.UnfollowUser(toBan, myName)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = rt.db.UnfollowUser(myName, toBan)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	return
}
