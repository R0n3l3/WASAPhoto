package database

import "database/sql"

func (db *appdbimpl) GetMyStream(u string) ([]Photo, error) {
	userId, _ := db.GetUserId(u)
	var photos []Photo
	res, err := db.c.Query("SELECT photoId FROM photos p WHERE EXISTS(SELECT * FROM follow WHERE follower=? AND following=p.uploader) ORDER BY uploadTime DESC", userId)
	if err != nil {
		panic(err)
	}
	defer func(res *sql.Rows) {
		err = res.Close()
		if err != nil {
			panic(err)
		}
	}(res)
	for res.Next() {
		var photo Photo
		err = res.Scan(&photo)
		if err != nil {
			panic(err)
		}
		photos = append(photos, photo)
	}
	return photos, nil
}
