package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func (db *appdbimpl) LikePhoto(photoId uint64, liker string) (Like, error) {
	var like Like

	id, err := db.GetUserId(liker)
	if err != nil {
		log.Println(err.Error())
		return like, err
	}

	result, err := db.c.Exec("SELECT likeId FROM likes l WHERE l.photoLiked=? AND l.liker=?", photoId, liker)
	if errors.Is(err, sql.ErrNoRows) {
		err = fmt.Errorf("you already liked this photo")
		log.Println("You already liked the photo")
		return like, err
	}

	result, err = db.c.Exec("UPDATE photos SET likeNumber=likeNumber+1 WHERE photoId=?", photoId)
	if err != nil {
		log.Println(err.Error())
		return like, err
	} else {
		rowsAffected, _ := result.RowsAffected()
		log.Println("Rows affected:", rowsAffected)
	}

	result, err = db.c.Exec("INSERT INTO likes(liker, photoLiked) VALUES (?, ?)", id, photoId)
	if err != nil {
		log.Println(err.Error())
		return like, err
	}
	likeId, err := result.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		return like, err
	}
	like.LikeId = uint64(likeId)
	like, err = db.GetLike(like.LikeId)
	if err != nil {
		log.Println(err.Error())
		return like, err
	}
	return like, nil
}
