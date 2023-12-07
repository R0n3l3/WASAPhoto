package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	uploaderName := ps.ByName("userId")           //Get the uploader name
	myName := r.URL.Query().Get("username")       //Get my name
	id, err := strconv.Atoi(ps.ByName("photoId")) //Get the photo id
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	comment, err := strconv.Atoi(r.URL.Query().Get("commentId")) //Get the comment id
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	uploaderProfile := getProfile(uploaderName) //Get the uploader name
	if uploaderProfile.ProfileId == "" {        //If null, it means that the user does not exist
		w.WriteHeader(http.StatusNotFound)
		return
	}
	myUser := getUser(myName)  //Get my user
	if myUser.Username == "" { //If null, it means that the user does not exist
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var photo Photo //Get the photo commented

	for j := 0; j <= len(uploaderProfile.Photos); j++ {
		if uploaderProfile.Photos[j].PhotoId == id {
			uploaderProfile.Photos[j].CommentNumber -= 1
			photo = uploaderProfile.Photos[j]
			break
		}
	}

	if photo.UploadTime == "" { //The photo couldn't be found
		w.WriteHeader(http.StatusNotFound)
		return
	}

	for k := 0; k < len(photo.Comments); k++ { //Remove the comment from the array and update the total
		if photo.Comments[k].CommentId == comment {
			if photo.Comments[k].Commenter.Username != myUser.Username { //If I didn't write the comment, error
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			photo.CommentNumber -= 1
			photo.Comments = append(photo.Comments[:k], photo.Comments[k+1:]...)
			break
		}
	}

	//Remove the comment from the general collection
	for i := 0; i < len(allComments); i++ {
		if allComments[i].CommentId == comment {
			allComments = append(allComments[:i], allComments[i+1:]...)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound) //If I didn't find the comment, error
	return
}
