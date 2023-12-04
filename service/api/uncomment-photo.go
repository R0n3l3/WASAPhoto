package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("userId")
	id, err := strconv.Atoi(ps.ByName("photoId"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	comment, err := strconv.Atoi(r.URL.Query().Get("commentId"))
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

	var photo Photo

	for j := 0; j <= len(profileMe.Photos); j++ {
		if profileMe.Photos[j].PhotoId == id {
			profileMe.Photos[j].CommentNumber -= 1
			photo = profileMe.Photos[j]
			break
		}
	}

	for k := 0; k < len(photo.Comments); k++ {
		if photo.Comments[k].CommentId == comment {
			photo.Comments = append(photo.Comments[:k], photo.Comments[k+1:]...)
			break
		}
	}
}
