package database

import (
	"database/sql"
	"errors"
	"log"
)

func (db *appdbimpl) GetComment(id uint64) (Comment, error) {
	var comment Comment

	err := db.c.QueryRow("SELECT commentId, commenter, commentTime, content, photoComment  FROM comments WHERE commentId= ?", id).Scan(&comment.CommentId, &comment.Commenter, &comment.CommentTime, &comment.Content, &comment.PhotoComment)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Println(err.Error())
		}
		return comment, err
	}
	return comment, nil
}
