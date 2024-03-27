package database

import (
	"database/sql"
	"errors"
	"log"
)

func (db *appdbimpl) GetMyPhotos(name string) ([]Photo, error) {
	var photos []Photo

	userId, err := db.GetUserId(name)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Println(err.Error())
		}
		return nil, err
	}

	res, err := db.c.Query("SELECT photoId, uploader, uploadTime, likeNumber, commentNumber, image FROM photos p WHERE p.uploader=? ORDER BY uploadTime DESC", userId)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Println(err.Error())
		}
		return nil, err
	}
	defer func(rows *sql.Rows) {
		if closeErr := res.Close(); closeErr != nil {
			log.Println(closeErr.Error())
		}
	}(res)
	for res.Next() {
		var photo Photo
		err = res.Scan(&photo.PhotoId, &photo.Uploader, &photo.UploadTime, &photo.LikeNumber, &photo.CommentNumber, &photo.Image)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		photos = append(photos, photo)
	}
	if err = res.Err(); err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return photos, nil
}
