package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	name := ps.ByName("userId")
	toUnBan := r.URL.Query().Get("")

	for i := 0; i <= len(Users); i++ {
		if Users[i].Username == name {
			for j := 0; j <= len(Users); j++ {
				if Users[j].Username == toUnBan {
					//todo
				}
			}
		}
	}

}
