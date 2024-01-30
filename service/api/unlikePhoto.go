package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	like, err := strconv.Atoi(r.URL.Query().Get("likeId")) //Get the like I want to remove
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = rt.db.UnlikePhoto(int64(like))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
