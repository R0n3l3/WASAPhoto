package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	me := ps.ByName("userId")              //Get my name
	toUnfollow := ps.ByName("followingId") //Get the unfollow name

	isAuth := rt.db.IsAuthorized(getToken(r.Header.Get("Authorization")))
	if !isAuth {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err := rt.db.UnfollowUser(toUnfollow, me)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

}
