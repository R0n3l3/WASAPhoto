package database

import (
	"database/sql"
	"errors"
	"log"
	"strings"
)

func (db *appdbimpl) LikePhoto(photoId uint64, liker string) (Like, error) {
	var like Like

	id, err := db.GetUserId(liker)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Println(err.Error())
		}
		return like, err
	}

	result, err := db.c.Exec("INSERT INTO likes(liker, photoLiked) VALUES (?, ?)", id, photoId)
	if err != nil {
		if !strings.Contains(err.Error(), "UNIQUE constraint failed") {
			log.Println(err.Error())
		}
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
}
