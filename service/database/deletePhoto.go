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
		if !errors.Is(err, sql.ErrNoRows) {
			log.Println(err.Error())
		}
		return err

	}
	profile, err := db.GetUserProfile(username)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Println(err.Error())
		}
		return err

	}
	if _, err = db.c.Exec("DELETE FROM likes WHERE photoLiked=?", id); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Println(err.Error())
		}
		return err
	}
	if _, err = db.c.Exec("DELETE FROM comments WHERE photoComment=?", id); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Println(err.Error())
		}
		return err
	}
	if _, err = db.c.Exec("DELETE FROM photos WHERE photoId=?", id); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Println(err.Error())
		}
		return err
	}

	_, err = db.c.Exec("UPDATE profiles SET photoNumber=profiles.photoNumber-1 WHERE profileId=?", profile.ProfileId)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Println(err.Error())
		}
		return err
	}
	return nil
}
