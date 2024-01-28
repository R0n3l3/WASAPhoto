package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func (db *appdbimpl) CommentPhoto(photoId int64, commenter string, content string) (int64, error) {
	id, _ := db.GetUserId(commenter)

	_, err := db.c.Exec("UPDATE photos SET commentNumber+=1 WHERE photoId=?", photoId)
	if err != nil {
		var photo Photo
		if err = db.c.QueryRow("SELECT photoId FROM photos WHERE photoId=?", photoId).Scan(&photo.PhotoId); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return photoId, fmt.Errorf("photo does not exist: %w", err)
			}
		}
	}
	res, err := db.c.Exec("INSERT INTO comments(commenter, commentTime, content, photoContent) VALUES (?, CURRENT_TIMESTAMP, ?, ?)", id, content, photoId)
	if err != nil {
		panic(err)
	}
	commentId, err := res.LastInsertId()
	if err != nil {
		var comment Comment
		if err := db.c.QueryRow("SELECT commentId FROM comments WHERE commentId=?", commentId).Scan(&comment.CommentId); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return commentId, fmt.Errorf("comment does not exist: %w", err)
			}
		}
	}
	return commentId, nil
}
