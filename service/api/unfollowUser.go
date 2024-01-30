package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	me := ps.ByName("userId")              //Get my name
	toUnfollow := ps.ByName("followingId") //Get the unfollow name

	err := rt.db.UnfollowUser(toUnfollow, me)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

}
