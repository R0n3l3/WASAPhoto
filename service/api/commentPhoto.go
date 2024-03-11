package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"strconv"
)

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("content-type", "application/json")
	uploaderName := ps.ByName("username")
	photo, err := strconv.Atoi(ps.ByName("photoId"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	text, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	isAuth := rt.db.IsAuthorized(getToken(r.Header.Get("Authorization")))
	if !isAuth {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	comment, err := rt.db.CommentPhoto(uint64(photo), uploaderName, string(text))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(comment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
