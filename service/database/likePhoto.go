package database

import "log"

func (db *appdbimpl) LikePhoto(photoId uint64, liker string) (Like, error) {
	var like Like

	id, err := db.GetUserId(liker)
	if err != nil {
		log.Println(err.Error())
		return like, err
	}

	_, err = db.c.Exec("UPDATE photos SET likeNumber=likeNumber+1 WHERE photoId=?", photoId)
	if err != nil {
		var photo Photo
		if err = db.c.QueryRow("SELECT photoId FROM photos WHERE photoId=?", photoId).Scan(&photo.PhotoId); err != nil {
			log.Println(err.Error())
			return like, err
		}
	}

	if err := db.c.QueryRow("INSERT INTO likes(liker, photoLiked) VALUES (?, ?)", id, photoId).Scan(like); err != nil {
		log.Println(err.Error())
		return like, err
	}
	return like, nil
}
