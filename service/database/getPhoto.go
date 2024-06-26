package database

import (
	"database/sql"
	"errors"
	"log"
)

func (db *appdbimpl) GetPhoto(id uint64) (Photo, error) {
	var photo Photo

	if err := db.c.QueryRow("SELECT photoId, uploadTime, commentNumber, likeNumber, uploader, image FROM photos WHERE photoId= ?", id).Scan(&photo.PhotoId, &photo.UploadTime, &photo.CommentNumber, &photo.LikeNumber, &photo.Uploader, &photo.Image); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Println(err.Error())
		}
		return photo, err
	}
	return photo, nil
}
