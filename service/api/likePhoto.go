package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")
	uploaderName := ps.ByName("userId")     //Get the name of the uploader of the photo
	myName := r.URL.Query().Get("username") //Get my name
	if myName == uploaderName {             //You cannot like your own photo!
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(ps.ByName("photoId")) //Get the id of the photo I want to like
	if err != nil {                               //If I get an error, stop
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	like, err := rt.db.LikePhoto(int64(id), myName)
	if err != nil {
		panic(err)
	}

	err = json.NewEncoder(w).Encode(like)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
