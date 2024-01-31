package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	comment, err := strconv.Atoi(r.URL.Query().Get("commentId")) //Get the comment id
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	isAuth := rt.db.IsAuthorized(getToken(r.Header.Get("Authorization")))
	if !isAuth {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = rt.db.UncommentPhoto(uint64(comment))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}
