package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	myName := ps.ByName("userId")    //Get my name
	toUnban := ps.ByName("bannedId") //Get the name of the unbanned user

	isAuth := rt.db.IsAuthorized(getToken(r.Header.Get("Authorization")))
	if !isAuth {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err := rt.db.UnbanUser(toUnban, myName)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}
