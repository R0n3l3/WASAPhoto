package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

var allLikes = []Like {

}

func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	name := ps.ByName("userId")
	photo, err := strconv.Atoi(ps.ByName("photoId"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var profileMe Profile

	for i:=0; i<len(Users); i++ {
		if Users[i].Username==name {
			profileMe=Users[i].UserProfile
			break
		}
	}

	newLike:= Like {
		LikeId: len(allLikes),
		Liker:
	}

	for i:=0; i<len(profileMe.Photos); i++ {
		if profileMe.Photos[i].PhotoId==photo {
			profileMe.Photos[i].Likes = append(profileMe.Photos[i].Likes, Like {

			})
		}
	}
}
