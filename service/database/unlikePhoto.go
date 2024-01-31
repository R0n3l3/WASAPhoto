package database

import (
	"database/sql"
	"errors"
	"log"
)

func (db *appdbimpl) UnlikePhoto(id uint64) error {
	var photoId uint64
	if err := db.c.QueryRow("SELECT photoLiked FROM likes WHERE likeId=?", id).Scan(&photoId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return err
		}
	}
	_, err := db.c.Exec("UPDATE photos SET likeNumber=likeNumber-1 WHERE photoId=?", photoId)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.c.Exec("DELETE FROM likes WHERE likeId=?", id)
	if err != nil {
		var like Like
		if errors.Is(db.c.QueryRow("SELECT likeID from likes WHERE likeId=?", id).Scan(&like.LikeId), sql.ErrNoRows) {
			return err
		}
		log.Fatal(err)
	}
	return nil
}
