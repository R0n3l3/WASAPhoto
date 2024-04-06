package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/R0n3l3/WASAPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("content-type", "application/json")

	isAuth := rt.db.IsAuthorized(getToken(r.Header.Get("Authorization")))
	if !isAuth {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	myName := ps.ByName("username")
	name := r.URL.Query().Get("search")

	id, err := strconv.ParseInt(name, 10, 64)
	if err != nil {
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
			w.WriteHeader(http.StatusNotFound)
			return
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
		w.WriteHeader(http.StatusOK)
	} else {
		profile, err := rt.db.GetUserProfileId(uint64(id))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}
		err = json.NewEncoder(w).Encode(profile)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)

	}
}
