package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"time"
)

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("content-type", "application/json")

	uploaderName := ps.ByName("userId")              // Get the uploader name
	myName := r.URL.Query().Get("username")          // Get my name
	photo, err := strconv.Atoi(ps.ByName("photoId")) // Get the commented photo
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	text := r.URL.Query().Get("content") // Get the content of the comment
	if text == "" {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	myUser := getUser(myName)  // Get my user
	if myUser.Username == "" { // If null, it means that the user does not exist
		w.WriteHeader(http.StatusNotFound)
		return
	}
	uploaderProfile := getProfile(uploaderName) // Get uploader profile
	if uploaderProfile.ProfileId == "" {        // If null, it means that the user does not exist
		w.WriteHeader(http.StatusNotFound)
		return
	}

	newComment := Comment{ // Create a new comment
		CommentId:   len(allComments),
		Commenter:   myUser,
		CommentTime: time.Now().Format("2006-01-02 15:04:51"),
		Content:     text,
	}

	allComments = append(allComments, newComment) // Add the comment to the general collection

	for i := 0; i < len(uploaderProfile.Photos); i++ { // Add the comment to the collection and update the total
		if uploaderProfile.Photos[i].PhotoId == photo {
			uploaderProfile.Photos[i].CommentNumber += 1
			uploaderProfile.Photos[i].Comments = append(uploaderProfile.Photos[i].Comments, newComment)
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
