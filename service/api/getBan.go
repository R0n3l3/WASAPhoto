package api

import (
	"database/sql"
	"errors"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

func (rt *_router) getBan(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	isAuth := rt.db.IsAuthorized(getToken(r.Header.Get("Authorization")))
	if !isAuth {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	banner := ps.ByName("username")
	banned, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	isBanned, err := rt.db.IsBanned(banner, string(banned))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
	}
	if isBanned {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
