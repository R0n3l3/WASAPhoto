package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"encoding/json"
)

type User struct {
	username	string
	banned 		[]string
	profile 	string
}

var Users = []User {
	User {
		username: "Maria"
		banned: ["Paolo", "Luigi"]
		profile: "Maria"
	},

	User {
		username: "Paolo"
		banned: ["Maria", "Giulia"]
		profile: "Paolo"
	}
}

type Profile struct {
	profileId 	string
	photos		[]int
	photoNumber	int
	followers 	[]string
	following	[]string
}

var Profiles = []Profile {
	Profile {
		profileId: "Maria"
		photos: [777, 431, 920]
		photoNumber: 3
		followers: ["Giulia", "Andrea"]
		following: ["Giulia", "Andrea", "Pino"]
	},
	
	Profile {
		profileId: "Paolo"
		photos: []
		photoNumber: 0
		followers: ["Andrea"]
		following: []
	}
}

// Obtain the user profile connected to the username
func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	searchUser := r.URL.Query().Get("username")
	
	if len(searchUser) < 3 || len(searchUser) > 16:
		w.WriteHeader(http.StatusNotAcceptable)
		return
	
	for i:=0; i<len(Users); i++ {
		if Users[i].username == searchUser:
			
	} 
}