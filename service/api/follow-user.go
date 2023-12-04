package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	toFollow := r.URL.Query().Get("username")
	me := ps.ByName("userId")

	var userFollow User
	var userMe User

	for i := 0; i < len(Users); i++ {
		if Users[i].Username == toFollow {
			userFollow = Users[i]
		}
		if Users[i].Username == me {
			userMe = Users[i]
		}
	}

	userMe.UserProfile.Following = append(userMe.UserProfile.Following, userFollow)
	userFollow.UserProfile.Followers = append(userFollow.UserProfile.Followers, userMe)
}
