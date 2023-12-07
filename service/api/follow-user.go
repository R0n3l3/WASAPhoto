package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	toFollow := r.URL.Query().Get("username") //Get the name I want to follow
	me := ps.ByName("userId")                 //Get my name

	userFollow := getUser(toFollow) //Get the user I want to follow
	if userFollow.Username == "" {  //If null, it means that the user does not exist
		w.WriteHeader(http.StatusNotFound)
		return
	}

	userMe := getUser(me)      //Get my user
	if userMe.Username == "" { //If null, it means that the user does not exist
		w.WriteHeader(http.StatusNotFound)
		return
	}

	userMe.UserProfile.Following = append(userMe.UserProfile.Following, userFollow)     //Add to follow to my following
	userFollow.UserProfile.Followers = append(userFollow.UserProfile.Followers, userMe) //Add me to their followers
	err := json.NewEncoder(w).Encode(userFollow.UserProfile)                            //Return the profile followed
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
