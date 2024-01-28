package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func (db *appdbimpl) DeletePhoto(id int64) error {
	var uploaderId int64
	var username string
	if err := db.c.QueryRow("SELECT uploader, username FROM photos, users WHERE photoId=? AND userId=uploader", id).Scan(&uploaderId, &username); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("no matching rows found: %w", err)
		}
	}
	uploaderId, _ = db.GetUserProfile(username)
	_, err := db.c.Exec("UPDATE profiles SET photoNumber-=1 WHERE profileId=?", uploaderId)
	if err != nil {
		panic(err)
	}

	_, err = db.c.Exec("DELETE FROM photos WHERE photoId=?", id)
	if err != nil {
		var photo Photo
		if errors.Is(db.c.QueryRow("SELECT photoId from photos WHERE photoId=?", id).Scan(&photo.PhotoId), sql.ErrNoRows) {
			return fmt.Errorf("no matching rows found: %w", err)
		}
	}

	return nil
}