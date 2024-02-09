package database

func (db *appdbimpl) UploadPhoto(uploader string, image []byte) (Photo, error) {
	var photo Photo

	id, err := db.GetUserId(uploader)
	if err != nil {
		print(err.Error())
		return photo, err
	}

	if err = db.c.QueryRow("INSERT INTO photos(uploader, uploadTime, likeNumber, commentNumber, image) VALUES(?, CURRENT_TIMESTAMP, 0, 0, ?)", id, image).Scan(&photo); err != nil {
		print(err.Error())
		return photo, err
	}

	profile, err := db.GetUserProfile(uploader)
	if err != nil {
		print(err.Error())
		return photo, err
	}

	_, err = db.c.Exec("UPDATE profiles SET photoNumber=profiles.photoNumber+1 WHERE profileId=?", profile.ProfileId)
	if err != nil {
		print(err.Error())
		return photo, err
	}

	return photo, nil
}
