package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	myName := ps.ByName("userId") //Recover my name

	myProfile := getProfile(myName) //Recover my profile

	if myProfile.ProfileId == "" { //If null, it means that the user does not exist
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var myStream []Photo //Recover the photos uploaded by the users I follow

	for i := 0; i < len(myProfile.Following); i++ {
		photos := myProfile.Following[i].UserProfile.Photos
		for j := 0; i < len(photos); j++ {
			myStream = append(myStream, photos[j])
		}
	}

	for i := 0; i < len(myStream); i++ { //Sort the photos by upload time
		for j := 0; j < len(myStream)-1; j++ {
			if myStream[j].UploadTime > myStream[j+1].UploadTime {
				myStream[j], myStream[j+1] = myStream[j+1], myStream[j]
			}
		}
	}

	err := json.NewEncoder(w).Encode(myStream) //Show the photos
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
