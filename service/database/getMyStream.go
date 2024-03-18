package database

import (
	"database/sql"
	"log"
)

func (db *appdbimpl) GetMyStream(u string) ([]Photo, error) {
	var stream []Photo

	userId, err := db.GetUserId(u)
	if err != nil {
		log.Println(err.Error())
		return stream, err
	}

	res, err := db.c.Query("SELECT p.photoId, p.uploader, p.uploadTime, p.likeNumber, p.commentNumber, p.image FROM photos p, follow f WHERE f.following=p.uploader AND f.follower= ? ORDER BY uploadTime DESC", userId)
	if err != nil {
		log.Println(err.Error())
		return stream, err
	}
	defer func(rows *sql.Rows) {
		if closeErr := res.Close(); closeErr != nil {
			log.Println(closeErr.Error())
		}
	}(res)
	for res.Next() {
		var photo Photo
		err = res.Scan(&photo.PhotoId, &photo.Uploader, &photo.UploadTime, &photo.LikeNumber, &photo.CommentNumber, &photo.Image)
		if err != nil {
			log.Println(err.Error())
			return stream, err
		}
		stream = append(stream, photo)
	}
	if err = res.Err(); err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return stream, nil
}
