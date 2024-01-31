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
		if errors.Is(err, sql.ErrNoRows) {
			return photo, err
		}
		log.Fatal(err)
	}

	if err = db.c.QueryRow("INSERT INTO photos(uploader, uploadTime, likeNumber, commentNumber, image) VALUES(?, CURRENT_TIMESTAMP, 0, 0, ?)", id, image).Scan(&photo); err != nil {
		log.Fatal(err)
	}

	profile, err := db.GetUserProfile(uploader)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return photo, err
		}
		log.Fatal(err)
	}

	_, err = db.c.Exec("UPDATE profiles SET photoNumber=profiles.photoNumber+1 WHERE profileId=?", profile.ProfileId)
	if err != nil {
		log.Fatal(err)
	}

	return photo, nil
}
