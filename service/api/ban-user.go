package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	name := ps.ByName("userId")
	toBan := r.URL.Query().Get("username")

	var userMe User
	var userToBan User

	for i := 0; i <= len(Users); i++ {
		if Users[i].Username == name {
			userMe = Users[i]
		}
		if Users[i].Username == toBan {
			userToBan = Users[i]
		}
	}
	userMe.Banned = append(userMe.Banned, userToBan)
}
