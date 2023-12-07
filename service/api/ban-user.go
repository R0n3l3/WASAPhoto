package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	myName := ps.ByName("userId")          //Get my name
	toBan := r.URL.Query().Get("username") //Get the name of the banned person

	myUser := getUser(myName)  //Get my user
	if myUser.Username == "" { //If null, it means that the user does not exist
		w.WriteHeader(http.StatusNotFound)
		return
	}
	toBanUser := getUser(toBan)   //Get the user to ban
	if toBanUser.Username == "" { //If null, it means that the user does not exist
		w.WriteHeader(http.StatusNotFound)
		return
	}

	//Add the banned person to the list, and remove them from the following and followers collection
	myUser.Banned = append(myUser.Banned, toBanUser)

	for i := 0; i < len(myUser.UserProfile.Followers); i++ {
		if myUser.UserProfile.Followers[i].Username == toBan {
			myUser.UserProfile.Followers = append(myUser.UserProfile.Followers[:i], myUser.UserProfile.Followers[i+1:]...)
		}
	}
	for i := 0; i < len(myUser.UserProfile.Following); i++ {
		if myUser.UserProfile.Following[i].Username == toBan {
			myUser.UserProfile.Following = append(myUser.UserProfile.Following[:i], myUser.UserProfile.Following[i+1:]...)
		}
	}
	for i := 0; i < len(toBanUser.UserProfile.Following); i++ {
		if toBanUser.UserProfile.Following[i].Username == myName {
			toBanUser.UserProfile.Following = append(toBanUser.UserProfile.Following[:i], toBanUser.UserProfile.Following[i+1:]...)
		}
	}
	for i := 0; i < len(toBanUser.UserProfile.Followers); i++ {
		if toBanUser.UserProfile.Followers[i].Username == myName {
			toBanUser.UserProfile.Followers = append(toBanUser.UserProfile.Followers[:i], toBanUser.UserProfile.Followers[i+1:]...)
		}
	}
	err := json.NewEncoder(w).Encode(myUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
