package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func (db *appdbimpl) LikePhoto(photoId int64, liker string) (int64, error) {
	id, _ := db.GetUserId(liker)

	_, err := db.c.Exec("UPDATE photos SET likeNumber=likeNumber+1 WHERE photoId=?", photoId)
	if err != nil {
		var photo Photo
		if err = db.c.QueryRow("SELECT photoId FROM photos WHERE photoId=?", photoId).Scan(&photo.PhotoId); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return photoId, fmt.Errorf("photo does not exist: %w", err)
			}
		}
	}
	res, err := db.c.Exec("INSERT INTO likes(liker, photoLiked) VALUES (?, ?)", id, photoId)
	if err != nil {
		panic(err)
	}
	likeId, err := res.LastInsertId()
	if err != nil {
		var like Like
		if err := db.c.QueryRow("SELECT likeId FROM likes WHERE likeId=?", likeId).Scan(&like.LikeId); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return likeId, fmt.Errorf("like does not exist: %w", err)
			}
		}
	}
	return likeId, nil
}
