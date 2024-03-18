package database

import (
	"log"
)

func (db *appdbimpl) DeletePhoto(id uint64) error {
	var uploaderId uint64
	var username string
	if err := db.c.QueryRow("SELECT uploader, username FROM photos, users WHERE photoId=? AND userId=uploader", id).Scan(&uploaderId, &username); err != nil {
		log.Println(1, err.Error())
		return err

	}

	profile, err := db.GetUserProfile(username)
	if err != nil {
		log.Println(2, err.Error())
		return err

	}

	res, err := db.c.Exec("UPDATE profiles SET photoNumber=profiles.photoNumber-1 WHERE profileId=?", profile.ProfileId)
	if err != nil {
		log.Println(3, err.Error())
		return err
	} else {
		rowsAffected, _ := res.RowsAffected()
		log.Println("Rows affected:", rowsAffected)
	}
	if _, err = db.c.Exec("DELETE FROM comments WHERE photoComment=?", id); err != nil {
		log.Println(6, err.Error())
		return err
	}
	if _, err = db.c.Exec("DELETE FROM likes WHERE photoLiked=?", id); err != nil {
		log.Println(5, err.Error())
		return err
	}

	if _, err = db.c.Exec("DELETE FROM photos WHERE photoId=?", id); err != nil {
		log.Println(4, err.Error())
		return err
	}

	return nil
}
