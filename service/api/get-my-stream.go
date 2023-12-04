package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	name := ps.ByName("userId")

	var profileMe Profile

	for i := 0; i < len(Users); i++ {
		if Users[i].Username == name {
			profileMe = Users[i].UserProfile
		}
	}

	var stream []Photo

	for i := 0; i < len(profileMe.Following); i++ {
		photos := profileMe.Following[i].UserProfile.Photos
		for j := 0; i < len(photos); j++ {
			stream = append(stream, photos[j])
		}
	}

	for i := 0; i < len(stream); i++ {
		for j := 0; j < len(stream)-1; j++ {
			if stream[j].UploadTime > stream[j+1].UploadTime {
				stream[j], stream[j+1] = stream[j+1], stream[j]
			}
		}
	}

	json.NewEncoder(w).Encode(stream)
}
