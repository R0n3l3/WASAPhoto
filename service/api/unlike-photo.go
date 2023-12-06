package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	uploaderName := ps.ByName("userId")     //Get the uploader of the photo
	myName := r.URL.Query().Get("username") //Get my username

	id, err := strconv.Atoi(ps.ByName("photoId")) //Get the id of the photo I want to unlike
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	like, err := strconv.Atoi(r.URL.Query().Get("likeId")) //Get the like I want to remove
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	uploaderProfile := getProfile(uploaderName) //Get the uploader profile
	if uploaderProfile.ProfileId == "" {        //If null, it means that the user does not exist
		w.WriteHeader(http.StatusNotFound)
		return
	}
	myUser := getUser(myName)  //Get my user
	if myUser.Username == "" { //If null, it means that the user does not exist
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var photo Photo //Get the photo I want to remove the like from

	for j := 0; j <= len(uploaderProfile.Photos); j++ {
		if uploaderProfile.Photos[j].PhotoId == id {
			uploaderProfile.Photos[j].LikeNumber -= 1
			photo = uploaderProfile.Photos[j]
			break
		}
	}

	if photo.UploadTime == "" { //If I didn't find the photo, return error
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if photo.Uploader.Username != myUser.Username { //If I didn't put the like, return error
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	for k := 0; k < len(photo.Likes); k++ { //Remove the like from the collection and update the total
		if photo.Likes[k].LikeId == like {
			photo.Likes = append(photo.Likes[:k], photo.Likes[k+1:]...)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound) //If I didn't find the like, return error
	return
}
