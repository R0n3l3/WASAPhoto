package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func (db *appdbimpl) UploadPhoto(uploader string, image []byte) (Photo, error) {
	id, _ := db.GetUserId(uploader)

	var photo Photo

	if err := db.c.QueryRow("INSERT INTO photos(uploader, uploadTime, likeNumber, commentNumber, image) VALUES(?, CURRENT_TIMESTAMP, 0, 0, ?)", id, image).Scan(&photo); err != nil {
		panic(err) //TODO
	}

	profile, _ := db.GetUserProfile(uploader)
	_, err := db.c.Exec("UPDATE profiles SET photoNumber=profiles.photoNumber+1 WHERE profileId=?", profile.ProfileId)
	if err != nil {
		var profile Profile
		if errors.Is(db.c.QueryRow("SELECT profileId, profileName FROM profiles WHERE profileName=?", uploader).Scan(&profile.ProfileId, &profile.ProfileName), sql.ErrNoRows) {
			return photo, fmt.Errorf("profile does not exist: %w", err)
		}
	}
	return photo, nil
}
