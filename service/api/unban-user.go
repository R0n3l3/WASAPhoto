package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	name := ps.ByName("userId")
	toUnban := ps.ByName("bannedId")

	var userMe User
	var index int

	for i := 0; i <= len(Users); i++ {
		if Users[i].Username == name {
			userMe = Users[i]
			break
		}
	}
	for i := 0; i <= len(Users); i++ {
		if userMe.Banned[i].Username == toUnban {
			index = i
			break
		}
	}
	userMe.Banned = append(userMe.Banned[:index], userMe.Banned[index+1:]...)
}
