package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	toFollow := r.URL.Query().Get("username") //Get the name I want to follow
	me := ps.ByName("userId")                 //Get my name

	isAuth := rt.db.IsAuthorized(getToken(r.Header.Get("Authorization")))
	if !isAuth {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	isBanned, err := rt.db.IsBanned(toFollow, me)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if isBanned {
		err := json.NewEncoder(w).Encode("Profile not found")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	user, err := rt.db.FollowUser(toFollow, me)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
