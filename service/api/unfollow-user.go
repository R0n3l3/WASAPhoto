package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	name := ps.ByName("userId")
	toUnfollow := ps.ByName("followingId")

	var profileMe Profile
	var profileUnfollow Profile
	var index int
	var index2 int

	for i := 0; i <= len(Users); i++ {
		if Users[i].Username == name {
			profileMe = Users[i].UserProfile
			break
		}
	}
	for i := 0; i <= len(profileMe.Following); i++ {
		if profileMe.Following[i].Username == toUnfollow {
			index = i
			profileUnfollow = profileMe.Following[i].UserProfile
			break
		}
	}

	for i := 0; i < len(profileUnfollow.Followers); i++ {
		if profileUnfollow.Followers[i].Username == name {
			index2 = i
			break
		}
	}

	profileMe.Following = append(profileMe.Following[:index], profileMe.Following[index+1:]...)
	profileUnfollow.Followers = append(profileUnfollow.Followers[:index2], profileUnfollow.Followers[index2+1:]...)
}
