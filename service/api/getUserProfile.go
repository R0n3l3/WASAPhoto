package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	myName := ps.ByName("username")
	name := r.URL.Query().Get("username")

	isAuth := rt.db.IsAuthorized(getToken(r.Header.Get("Authorization")))
	if !isAuth {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	isBanned, err := rt.db.IsBanned(name, myName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if isBanned {
		err := json.NewEncoder(w).Encode("Profile not found")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	profile, err := rt.db.GetUserProfile(name)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	err = json.NewEncoder(w).Encode(profile)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
