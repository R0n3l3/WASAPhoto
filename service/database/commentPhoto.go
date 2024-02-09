package database

func (db *appdbimpl) CommentPhoto(photoId uint64, commenter string, content string) (Comment, error) {
	var comment Comment

	id, err := db.GetUserId(commenter)
	if err != nil {
		print(err.Error())
		return comment, err
	}

	_, err = db.c.Exec("UPDATE photos SET commentNumber=commentNumber+1 WHERE photoId=?", photoId)
	if err != nil {
		var photo Photo
		if err = db.c.QueryRow("SELECT photoId FROM photos WHERE photoId=?", photoId).Scan(&photo.PhotoId); err != nil {
			print(err.Error())
			return comment, err
		}
	}
	if err := db.c.QueryRow("INSERT INTO comments(commenter, commentTime, content, photoComment) VALUES (?, CURRENT_TIMESTAMP, ?, ?)", id, content, photoId).Scan(&comment); err != nil {
		print(err.Error())
		return comment, err
	}

	return comment, nil
}
