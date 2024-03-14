package database

import "log"

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
	comment.CommentId = uint64(commentId)
	comment, err = db.GetComment(comment.CommentId)
	if err != nil {
		log.Println(err.Error())
		return comment, err
	}

	result, err = db.c.Exec("UPDATE photos SET commentNumber=commentNumber+1 WHERE photoId=?", photoId)
	if err != nil {
		log.Println(err.Error())
		return comment, err
	} else {
		rowsAffected, _ := result.RowsAffected()
		log.Println("Rows affected:", rowsAffected)
	}

	return comment, nil
}
