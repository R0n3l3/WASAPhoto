package database

import (
	"database/sql"
	"errors"
	"log"
)

func (db *appdbimpl) UncommentPhoto(commentId uint64, photoId uint64) error {
	_, err := db.c.Exec("DELETE FROM comments WHERE commentId=?", commentId)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Println(err.Error())
		}
		return err
	}
	_, err = db.c.Exec("UPDATE photos SET commentNumber=commentNumber-1 WHERE photoId=?", photoId)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Println(err.Error())
		}
		return err
	}
	return nil
}
