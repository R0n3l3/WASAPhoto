package database

import (
	"database/sql"
	"errors"
	"log"
)

func (db *appdbimpl) CommentPhoto(photoId uint64, commenter string, content string) (Comment, error) {
	var comment Comment

	result, err := db.c.Exec("INSERT INTO comments(commenter, commentTime, content, photoComment) VALUES (?, CURRENT_TIMESTAMP, ?, ?)", commenter, content, photoId)
	if err != nil {
		log.Println(err.Error())
		return comment, err
	}

	commentId, err := result.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		return comment, err
	}

	comment, err = db.GetComment(uint64(commentId))
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Println(err.Error())
		}
		return comment, err
	}

	_, err = db.c.Exec("UPDATE photos SET commentNumber=commentNumber+1 WHERE photoId=?", photoId)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Println(err.Error())
		}
		return comment, err

	}
	return comment, nil
}
