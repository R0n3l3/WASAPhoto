package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	link := []byte(r.URL.Query().Get("link")) //Get the link of the image
	uploaderUsername := ps.ByName("userId")   //Get the name of the uploader

	isAuth := rt.db.IsAuthorized(getToken(r.Header.Get("Authorization")))
	if !isAuth {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	photo, _ := rt.db.UploadPhoto(uploaderUsername, link)

	err := json.NewEncoder(w).Encode(photo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
