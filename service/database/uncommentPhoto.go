package database

import "log"

func (db *appdbimpl) UncommentPhoto(commentId uint64) error {
	var photoId uint64
	if err := db.c.QueryRow("SELECT photoComment FROM comments WHERE commentId=?", commentId).Scan(&photoId); err != nil {
		log.Println(err.Error())
		return err
	}
	_, err := db.c.Exec("UPDATE photos SET commentNumber=commentNumber-1 WHERE photoId=?", photoId)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	_, err = db.c.Exec("DELETE FROM comments WHERE commentId=?", commentId)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
