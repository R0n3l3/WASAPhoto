package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	name := ps.ByName("userId")
	id, err := strconv.Atoi(ps.ByName("photoId"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	like, err := strconv.Atoi(r.URL.Query().Get("likeId"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var otherProfile Profile

	for i := 0; i <= len(Users); i++ {
		if Users[i].Username == name {
			otherProfile = Users[i].UserProfile
			break
		}
	}

	var photo Photo

	for j := 0; j <= len(otherProfile.Photos); j++ {
		if otherProfile.Photos[j].PhotoId == id {
			otherProfile.Photos[j].LikeNumber -= 1
			photo = otherProfile.Photos[j]
			break
		}
	}

	for k := 0; k < len(photo.Likes); k++ {
		if photo.Likes[k].LikeId == like {
			photo.Likes = append(photo.Likes[:k], photo.Likes[k+1:]...)
			break
		}
	}
}
