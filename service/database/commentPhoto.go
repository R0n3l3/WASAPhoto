package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func (db *appdbimpl) CommentPhoto(photoId int64, commenter string, content string) (Comment, error) {
	id, _ := db.GetUserId(commenter)
	var comment Comment

	_, err := db.c.Exec("UPDATE photos SET commentNumber=commentNumber+1 WHERE photoId=?", photoId)
	if err != nil {
		var photo Photo
		if err = db.c.QueryRow("SELECT photoId FROM photos WHERE photoId=?", photoId).Scan(&photo.PhotoId); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return comment, fmt.Errorf("photo does not exist: %w", err)
			}
		}
	}
	if err := db.c.QueryRow("INSERT INTO comments(commenter, commentTime, content, photoComment) VALUES (?, CURRENT_TIMESTAMP, ?, ?)", id, content, photoId).Scan(&comment); err != nil {
		panic(err)
	}

	return comment, nil
}
