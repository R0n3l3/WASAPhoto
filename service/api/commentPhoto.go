package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("content-type", "application/json")
	uploaderName := ps.ByName("userId")              //Get the uploader name
	photo, err := strconv.Atoi(ps.ByName("photoId")) //Get the commented photo
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	text := r.URL.Query().Get("content") //Get the content of the comment
	if text == "" {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

isAuth := rt.db.IsAuth(getToken(r.Header.Get("Authorization")))
	if !isAuth {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	
	comment, err := rt.db.CommentPhoto(uint64(photo), uploaderName, text)
	if err != nil {
		panic(err)
	}

	err = json.NewEncoder(w).Encode(comment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
