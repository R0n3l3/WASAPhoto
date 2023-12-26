package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	link := r.URL.Query().Get("link")       // Get the link of the image
	uploaderUsername := ps.ByName("userId") // Get the name of the uploader

	uploader := getUser(uploaderUsername) // Get the user of the uploader
	if uploader.Username == "" {          // If null, it means that the user does not exist
		w.WriteHeader(http.StatusNotFound)
		return
	}

	newPhoto := Photo{ // Create a new photo
		PhotoId:  len(allPhotos),
		Uploader: uploader,
		PhotoImage: Image{
			Link: link,
		},
		UploadTime:    time.Now().Format("2006-01-02 15:04:51"),
		Likes:         []Like{},
		LikeNumber:    0,
		Comments:      []Comment{},
		CommentNumber: 0,
	}

	allPhotos = append(allPhotos, newPhoto) // Add the photo to the general collection

	profile := uploader.UserProfile

	profile.PhotoNumber += 1                          // Update the total of photos
	profile.Photos = append(profile.Photos, newPhoto) // Add the photo to the collection
	err := json.NewEncoder(w).Encode(newPhoto)        // Show the new photo
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for i := 0; i < len(profile.Photos); i++ { // Sort the photos
		for j := 0; j < len(profile.Photos)-1; j++ {
			if profile.Photos[j].UploadTime > profile.Photos[j+1].UploadTime {
				profile.Photos[j], profile.Photos[j+1] = profile.Photos[j+1], profile.Photos[j]
			}
		}
	}
}
