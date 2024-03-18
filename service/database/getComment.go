package database

import "log"

func (db *appdbimpl) GetComment(id uint64) (Comment, error) {
	var comment Comment

	if err := db.c.QueryRow("SELECT commentId, commenter, commentTime, content, photoComment  FROM comments WHERE commentId= ?", id).Scan(&comment.CommentId, &comment.Commenter, &comment.CommentTime, &comment.Content, &comment.PhotoComment); err != nil {
		log.Println(err.Error())
		return comment, err
	}
	return comment, nil
}
