package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	myName := ps.ByName("userId")    // Get my name
	toUnban := ps.ByName("bannedId") // Get the name of the unbanned user

	myUser := getUser(myName)  // Get my user
	if myUser.Username == "" { // If null, it means that the user does not exist
		w.WriteHeader(http.StatusNotFound)
		return
	}

	for i := 0; i <= len(Users); i++ { // Remove the user from the collection
		if myUser.Banned[i].Username == toUnban {
			myUser.Banned = append(myUser.Banned[:i], myUser.Banned[i+1:]...)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound) // The user was not banned
}
