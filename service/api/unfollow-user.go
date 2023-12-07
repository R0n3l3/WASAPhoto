package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	me := ps.ByName("userId")              //Get my name
	toUnfollow := ps.ByName("followingId") //Get the unfollow name

	profileMe := getProfile(me)    //Get my profile
	if profileMe.ProfileId == "" { //If null, it means that the user does not exist
		w.WriteHeader(http.StatusNotFound)
		return
	}

	profileUnfollow := getProfile(toUnfollow) //Get their profile
	if profileUnfollow.ProfileId == "" {      //If null, it means that the user does not exist
		w.WriteHeader(http.StatusNotFound)
		return
	}

	//Remove the user from my following
	for i := 0; i <= len(profileMe.Following); i++ {
		if profileMe.Following[i].Username == toUnfollow {
			profileMe.Following = append(profileMe.Following[:i], profileMe.Following[i+1:]...)
			break
		}
	}

	//Remove me from followers
	for i := 0; i < len(profileUnfollow.Followers); i++ {
		if profileUnfollow.Followers[i].Username == me {
			profileUnfollow.Followers = append(profileUnfollow.Followers[:i], profileUnfollow.Followers[i+1:]...)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound) //If I couldn't find the user, error
	return
}
