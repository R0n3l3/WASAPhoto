package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	name := ps.ByName("userId")                   //Get my name
	id, err := strconv.Atoi(ps.ByName("photoId")) //Get the photo id
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	myProfile := getProfile(name)
	if myProfile.ProfileId == "" { //If null, it means that the user does not exist
		w.WriteHeader(http.StatusNotFound)
		return
	}

	//Update the total of photos and the collection
	myProfile.PhotoNumber -= 1
	for j := 0; j <= len(myProfile.Photos); j++ {
		if myProfile.Photos[j].PhotoId == id {
			myProfile.Photos = append(myProfile.Photos[:j], myProfile.Photos[j+1:]...)
			break
		}
	}

	for i := 0; i < len(allPhotos); i++ { //Remove the photo from the general collection
		if allPhotos[i].PhotoId == id {
			allPhotos = append(allPhotos[:i], allPhotos[i+1:]...)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound) //If I didn't find the photo, return error
	return
}
