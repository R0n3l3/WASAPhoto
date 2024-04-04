package database

import (
	"database/sql"
	"errors"
	"log"
)

func (db *appdbimpl) DeletePhoto(id uint64, username string) error {
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
