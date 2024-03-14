package database

import (
	"database/sql"
	"errors"
	"log"
)

func (db *appdbimpl) DeletePhoto(id uint64) error {
	var uploaderId uint64
	var username string
	if err := db.c.QueryRow("SELECT uploader, username FROM photos, users WHERE photoId=? AND userId=uploader", id).Scan(&uploaderId, &username); err != nil {
		log.Println(err.Error())
		return err

	}

	profile, err := db.GetUserProfile(username)
	if err != nil {
		log.Println(err.Error())
		return err

	}

	res, err := db.c.Exec("UPDATE profiles SET photoNumber=profiles.photoNumber-1 WHERE profileId=?", profile.ProfileId)
	if err != nil {
		log.Println(err.Error())
		return err
	} else {
		rowsAffected, _ := res.RowsAffected()
		log.Println("Rows affected:", rowsAffected)
	}

	_, err = db.c.Exec("DELETE FROM photos WHERE photoId=?", id)
	if err != nil {
		var photo Photo
		if errors.Is(db.c.QueryRow("SELECT photoId from photos WHERE photoId=?", id).Scan(&photo.PhotoId), sql.ErrNoRows) {
			log.Println(err.Error())
			return err
		}
	}

	_, err = db.c.Exec("DELETE FROM likes WHERE photoLiked=?", id)
	if err != nil {
		var photo Photo
		if errors.Is(db.c.QueryRow("SELECT photoId from photos WHERE photoId=?", id).Scan(&photo.PhotoId), sql.ErrNoRows) {
			log.Println(err.Error())
			return err
		}
	}

	_, err = db.c.Exec("DELETE FROM comments WHERE photoComment=?", id)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
