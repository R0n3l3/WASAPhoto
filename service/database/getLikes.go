package database

import (
	"database/sql"
	"log"
)

func (db *appdbimpl) GetLikes(id uint64) ([]Like, error) {
	var likes []Like

	res, err := db.c.Query("SELECT likeId, liker, photoLiked FROM likes l WHERE l.photoLiked=?", id)
	if err != nil {
		log.Println(err.Error())
		return likes, err
	}
	defer func(rows *sql.Rows) {
		if closeErr := res.Close(); closeErr != nil {
			log.Println(closeErr.Error())
		}
	}(res)
	for res.Next() {
		var like Like
		err = res.Scan(&like.LikeId, &like.Liker, &like.PhotoLiked)
		if err != nil {
			log.Println(err.Error())
			return likes, err
		}
		likes = append(likes, like)
	}
	if err = res.Err(); err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return likes, nil
}
