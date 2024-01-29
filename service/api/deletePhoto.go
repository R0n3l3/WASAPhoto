package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	id, err := strconv.Atoi(ps.ByName("photoId")) //Get the photo id
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = rt.db.DeletePhoto(int64(id))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

}
