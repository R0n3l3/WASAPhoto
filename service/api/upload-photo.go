package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

var allPhotos = []Photo{
	{
		PhotoId:  0,
		Uploader: Users[0],
		PhotoImage: Image{
			Link: "https://www.piggypet.it/wp-content/uploads/2023/08/Progetto-senza-titolo-2023-08-05T132259.115-870x563@2x.jpg",
		},
		UploadTime:    "2020-11-10T17:05:50Z",
		Likes:         []Like{},
		LikeNumber:    0,
		Comments:      []Comment{},
		CommentNumber: 0,
	},
}

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	link := r.URL.Query().Get("link")
	uploaderUsername := ps.ByName("userId")

	var uploader User

	for i := 0; i < len(Users); i++ {
		if Users[i].Username == uploaderUsername {
			uploader = Users[i]
			break
		}
	}

	newPhoto := Photo{
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

	allPhotos = append(allPhotos, newPhoto)

	uploader.UserProfile.PhotoNumber += 1
	uploader.UserProfile.Photos = append(uploader.UserProfile.Photos, newPhoto)
}
