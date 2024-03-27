package database

import (
	"database/sql"
	"errors"
	"log"
)

func (db *appdbimpl) UploadPhoto(uploader string, image []byte) (Photo, error) {
	var photo Photo

	id, err := db.GetUserId(uploader)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Println(err.Error())
		}
		return photo, err
	}

	result, err := db.c.Exec("INSERT INTO photos(uploader, uploadTime, likeNumber, commentNumber, image) VALUES(?, CURRENT_TIMESTAMP, 0, 0, ?)", id, image)
	if err != nil {
		log.Println(err.Error())
		return photo, err
	}
	photoId, err := result.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		return photo, err
	}
	photo, err = db.GetPhoto(uint64(photoId))
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Println(err.Error())
		}
		return photo, err
	}

	profile, err := db.GetUserProfile(uploader)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Println(err.Error())
		}
		return photo, err
	}

	_, err = db.c.Exec("UPDATE profiles SET photoNumber=profiles.photoNumber+1 WHERE profileId=?", profile.ProfileId)
	if err != nil {
		log.Println(err.Error())
		return photo, err
	}

	return photo, nil
}
