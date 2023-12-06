package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	uploaderName := ps.ByName("userId")     //Get the name of the uploader of the photo
	myName := r.URL.Query().Get("username") //Get my name
	if myName == uploaderName {             //You cannot like your own photo!
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(ps.ByName("photoId")) //Get the id of the photo I want to like
	if err != nil {                               //If I get an error, stop
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	myUser := getUser(myName)  //Get my User
	if myUser.Username == "" { //If null, it means that the user does not exist
		w.WriteHeader(http.StatusNotFound)
		return
	}

	uploaderProfile := getProfile(uploaderName) //Get the uploader profile
	if uploaderProfile.ProfileId == "" {        //If null, it means that the user does not exist
		w.WriteHeader(http.StatusNotFound)
		return
	}

	newLike := Like{ //Create a new instance of a like
		LikeId: len(allLikes),
		Liker:  myUser,
	}

	for i := 0; i < len(uploaderProfile.Photos); i++ { //Update the total number of likes and add that instance of like to the collection
		if uploaderProfile.Photos[i].PhotoId == id {
			uploaderProfile.Photos[i].LikeNumber += 1
			uploaderProfile.Photos[i].Likes = append(uploaderProfile.Photos[i].Likes, newLike)
			err := json.NewEncoder(w).Encode(uploaderProfile.Photos[i])
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}
	}
	w.WriteHeader(http.StatusNotFound)
	return
}
