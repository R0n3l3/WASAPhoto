package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) getFollowers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("content-type", "application/json")
	isAuth := rt.db.IsAuthorized(getToken(r.Header.Get("Authorization")))
	if !isAuth {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	me := ps.ByName("username")
	getFollowers := r.URL.Query().Get("followers")

	if getFollowers == "true" {
		followers, err := rt.db.GetFollowers(me)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = json.NewEncoder(w).Encode(followers)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	} else {
		following, err := rt.db.GetFollowing(me)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = json.NewEncoder(w).Encode(following)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}