package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var Users = []User{
	{
		Username:    "Eleonora",
		Banned:      []User{},
		UserProfile: Profiles[0],
	},
}

var Profiles = []Profile{
	{
		ProfileId:   "Eleonora",
		Photos:      []Photo{},
		PhotoNumber: 0,
		Followers:   []User{},
		Following:   []User{},
	},
}

// Create new user
func (rt *_router) login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	name := r.URL.Query().Get("username")

	for i := 0; i <= len(Users); i++ {
		if Users[i].Username == name {
			json.NewEncoder(w).Encode(name)
			return
		}
	}
	Profiles = append(Profiles, Profile{
		ProfileId:   name,
		Photos:      []Photo{},
		PhotoNumber: 0,
		Followers:   []User{},
		Following:   []User{},
	})
	Users = append(Users, User{
		Username:    name,
		Banned:      []User{},
		UserProfile: Profiles[len(Profiles)-1],
	})
	json.NewEncoder(w).Encode(name)
}
