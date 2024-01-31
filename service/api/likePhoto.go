package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")
	myName := r.URL.Query().Get("username") //Get my name

	id, err := strconv.Atoi(ps.ByName("photoId")) //Get the id of the photo I want to like
	if err != nil {                               //If I get an error, stop
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	isAuth := rt.db.IsAuthorized(getToken(r.Header.Get("Authorization")))
	if !isAuth {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	like, err := rt.db.LikePhoto(uint64(id), myName)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(like)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
