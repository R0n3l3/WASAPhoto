package database

import (
	"database/sql"
	"errors"
	"log"
)

func (db *appdbimpl) GetLike(id uint64) (Like, error) {
	var like Like

	err := db.c.QueryRow("SELECT likeId, liker, photoLiked  FROM likes WHERE likeId= ?", id).Scan(&like.LikeId, &like.Liker, &like.PhotoLiked)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Println(err.Error())
		}
		return like, err
	}
	return like, nil
}
