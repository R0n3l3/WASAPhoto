package database

import "log"

func (db *appdbimpl) GetLike(id uint64) (Like, error) {
	var like Like

	if err := db.c.QueryRow("SELECT likeId, liker, photoLiked  FROM likes WHERE likeId= ?", id).Scan(&like.LikeId, &like.Liker, &like.PhotoLiked); err != nil {
		log.Println(err.Error())
		return like, err
	}
	return like, nil
}
