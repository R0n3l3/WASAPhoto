package database

import (
	"database/sql"
	"errors"
	"log"
)

func (db *appdbimpl) LikePhoto(photoId uint64, liker string) (Like, error) {
	var like Like
	var exist bool

	id, err := db.GetUserId(liker)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Println(err.Error())
		}
		return like, err
	}

	err = db.c.QueryRow("SELECT 1 FROM likes WHERE liker=? AND photoLiked=?", id, photoId).Scan(&exist)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			result, err := db.c.Exec("INSERT INTO likes(liker, photoLiked) VALUES (?, ?)", id, photoId)
			if err != nil {
				log.Println(err.Error())
				return like, err
			}
			likeId, err := result.LastInsertId()
			if err != nil {
				log.Println(err.Error())
				return like, err
			}
			like, err = db.GetLike(uint64(likeId))
			if err != nil {
				if !errors.Is(err, sql.ErrNoRows) {
					log.Println(err.Error())
				}
				return like, err
			}
			_, err = db.c.Exec("UPDATE photos SET likeNumber=likeNumber+1 WHERE photoId=?", photoId)
			if err != nil {
				if !errors.Is(err, sql.ErrNoRows) {
					log.Println(err.Error())
				}
				return like, err
			}
			return like, nil
		} else {
			log.Println(err)
			return like, err
		}
	}
	return like, nil
}
