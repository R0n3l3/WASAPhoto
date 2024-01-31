package database

import (
	"database/sql"
	"errors"
	"log"
)

func (db *appdbimpl) GetMyStream(u string) ([]Photo, error) {
	var photos []Photo

	userId, err := db.GetUserId(u)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return photos, err
		}
		log.Fatal(err)
	}

	res, err := db.c.Query("SELECT photoId FROM photos p WHERE EXISTS(SELECT * FROM follow WHERE follower=? AND following=p.uploader) ORDER BY uploadTime DESC", userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return photos, err
		}
		log.Fatal(err)
	}
	defer func(res *sql.Rows) {
		err = res.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res)
	for res.Next() {
		var photo Photo
		err = res.Scan(&photo)
		if err != nil {
			log.Fatal(err)
		}
		photos = append(photos, photo)
	}
	return photos, nil
}
