package database

import (
	"database/sql"
	"errors"
	"log"
)

func (db *appdbimpl) CommentPhoto(photoId uint64, commenter string, content string) (Comment, error) {
	var comment Comment

	id, err := db.GetUserId(commenter)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return comment, err
		}
		log.Fatal(err)
	}

	_, err = db.c.Exec("UPDATE photos SET commentNumber=commentNumber+1 WHERE photoId=?", photoId)
	if err != nil {
		var photo Photo
		if err = db.c.QueryRow("SELECT photoId FROM photos WHERE photoId=?", photoId).Scan(&photo.PhotoId); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return comment, err
			}
			log.Fatal(err)
		}
	}
	if err := db.c.QueryRow("INSERT INTO comments(commenter, commentTime, content, photoComment) VALUES (?, CURRENT_TIMESTAMP, ?, ?)", id, content, photoId).Scan(&comment); err != nil {
		log.Fatal(err)
	}

	return comment, nil
}
