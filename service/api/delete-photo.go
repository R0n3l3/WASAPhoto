package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	name := ps.ByName("Username")
	id, err := strconv.Atoi(ps.ByName("photoId"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var profileMe Profile

	for i := 0; i <= len(Users); i++ {
		if Users[i].Username == name {
			profileMe = Users[i].UserProfile
			break
		}
	}

	profileMe.PhotoNumber -= 1
	for j := 0; j <= len(profileMe.Photos); j++ {
		if profileMe.Photos[j].PhotoId == id {
			profileMe.Photos = append(profileMe.Photos[:j], profileMe.Photos[j+1:]...)
			break
		}
	}
}
