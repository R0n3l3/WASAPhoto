package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func (db *appdbimpl) LikePhoto(photoId int64, liker string) (Like, error) {
	id, _ := db.GetUserId(liker)
	var like Like

	_, err := db.c.Exec("UPDATE photos SET likeNumber=likeNumber+1 WHERE photoId=?", photoId)
	if err != nil {
		var photo Photo
		if err = db.c.QueryRow("SELECT photoId FROM photos WHERE photoId=?", photoId).Scan(&photo.PhotoId); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return like, fmt.Errorf("photo does not exist: %w", err)
			}
		}
	}
	if err := db.c.QueryRow("INSERT INTO likes(liker, photoLiked) VALUES (?, ?)", id, photoId).Scan(like); err != nil {
		panic(err)
	}
	return like, nil
}
