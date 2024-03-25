package database

import (
	"database/sql"
	"errors"
	"log"
)

func (db *appdbimpl) UnlikePhoto(id uint64, photoId uint64) error {
	_, err := db.c.Exec("DELETE FROM likes WHERE likeId=?", id)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Println(err.Error())
		}
		return err
	}
	_, err = db.c.Exec("UPDATE photos SET likeNumber=likeNumber-1 WHERE photoId=?", photoId)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Println(err.Error())
		}
		return err
	}
	return nil
}
