package database

import (
	"database/sql"
	"errors"
	"log"
)

func (db *appdbimpl) GetComments(id uint64) ([]Comment, error) {
	var comments []Comment

	res, err := db.c.Query("SELECT commentId, commenter, commentTime, content, photoComment FROM comments WHERE photoComment=? ORDER BY commentTime DESC", id)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Println(err.Error())
		}
		return nil, err
	}
	defer func(rows *sql.Rows) {
		if closeErr := res.Close(); closeErr != nil {
			log.Println(closeErr.Error())
		}
	}(res)

	for res.Next() {
		var comment Comment
		err = res.Scan(&comment.CommentId, &comment.Commenter, &comment.CommentTime, &comment.Content, &comment.PhotoComment)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		comments = append(comments, comment)
	}
	if err = res.Err(); err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return comments, nil
}
