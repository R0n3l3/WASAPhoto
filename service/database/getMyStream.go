package database

import (
	"database/sql"
	"log"
)

func (db *appdbimpl) GetMyStream(u string) ([]Photo, error) {
	var photos []Photo

	userId, err := db.GetUserId(u)
	if err != nil {
		log.Println(err.Error())
		return photos, err
	}

	res, err := db.c.Query("SELECT photoId FROM photos p WHERE EXISTS(SELECT * FROM follow WHERE follower=? AND following=p.uploader) ORDER BY uploadTime DESC", userId)
	if err != nil {
		log.Println(err.Error())
		return photos, err
	}
	func(res *sql.Rows) {
		err = res.Close()
		if err != nil {
			log.Println(err.Error())
			return
		}
	}(res)
	err = res.Err()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	for res.Next() {
		var photo Photo
		err = res.Scan(&photo)
		if err != nil {
			log.Println(err.Error())
			return photos, err
		}
		photos = append(photos, photo)
	}
	return photos, nil
}
