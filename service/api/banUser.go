package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	myName := ps.ByName("username")
	toBan, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	isAuth := rt.db.IsAuthorized(getToken(r.Header.Get("Authorization")))
	if !isAuth {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = rt.db.BanUser(string(toBan), myName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode("Ban successful")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return

	}
	err = rt.db.UnfollowUser(string(toBan), myName)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = rt.db.UnfollowUser(myName, string(toBan))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
