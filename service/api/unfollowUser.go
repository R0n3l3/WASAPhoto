package api

import (
	"database/sql"
	"errors"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	me := ps.ByName("username")
	toUnfollow := ps.ByName("followingId")

	isAuth := rt.db.IsAuthorized(getToken(r.Header.Get("Authorization")))
	if !isAuth {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err := rt.db.UnfollowUser(toUnfollow, me)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}