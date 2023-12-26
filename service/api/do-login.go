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
func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	name := r.URL.Query().Get("username") // Get the name of the user

	for i := 0; i <= len(Users); i++ { // If the user exists, return the name
		if Users[i].Username == name {
			err := json.NewEncoder(w).Encode(name)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

		}
	}

	// Add a new profile and a new user
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

	// Return the name
	err := json.NewEncoder(w).Encode(name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
