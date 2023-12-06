package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

var allLikes = []Like{}

func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	name := ps.ByName("userId")
	myName := r.URL.Query().Get("username")
	photo, err := strconv.Atoi(ps.ByName("photoId"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var userMe User
	var otherProfile Profile

	for i := 0; i < len(Users); i++ {
		if Users[i].Username == myName {
			userMe = Users[i]
			break
		}
		if Users[i].Username == name {
			otherProfile = Users[i].UserProfile
			break
		}
	}

	newLike := Like{
		LikeId: len(allLikes),
		Liker:  userMe,
	}

	for i := 0; i < len(otherProfile.Photos); i++ {
		if otherProfile.Photos[i].PhotoId == photo {
			otherProfile.Photos[i].LikeNumber += 1
			otherProfile.Photos[i].Likes = append(otherProfile.Photos[i].Likes, newLike)
			break
		}
	}
}
