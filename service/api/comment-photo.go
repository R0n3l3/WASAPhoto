package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"time"
)

var allComments = []Comment{}

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("content-type", "application/json")

	name := ps.ByName("userId")
	myName := r.URL.Query().Get("username")
	photo, err := strconv.Atoi(ps.ByName("photoId"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	text := r.URL.Query().Get("content")

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

	newComment := Comment{
		CommentId:   len(allComments),
		Commenter:   userMe,
		CommentTime: time.Now().Format("2006-01-02 15:04:51"),
		Content:     text,
	}

	for i := 0; i < len(otherProfile.Photos); i++ {
		if otherProfile.Photos[i].PhotoId == photo {
			otherProfile.Photos[i].CommentNumber += 1
			otherProfile.Photos[i].Comments = append(otherProfile.Photos[i].Comments, newComment)
			break
		}
	}
}
