package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func (db *appdbimpl) UncommentPhoto(commentId uint64) error {
	var photoId uint64
	if err := db.c.QueryRow("SELECT photoComment FROM comments WHERE commentId=?", commentId).Scan(&photoId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("no matching rows found: %w", err)
		}
	}
	_, err := db.c.Exec("UPDATE photos SET commentNumber=commentNumber-1 WHERE photoId=?", photoId)
	if err != nil {
		panic(err)
	}

	_, err = db.c.Exec("DELETE FROM comments WHERE commentId=?", commentId)
	if err != nil {
		var comment Comment
		if errors.Is(db.c.QueryRow("SELECT commentId from comments WHERE commentId=?", commentId).Scan(&comment.CommentId), sql.ErrNoRows) {
			return fmt.Errorf("no matching rows found: %w", err)
		}
	}
	return nil
}
