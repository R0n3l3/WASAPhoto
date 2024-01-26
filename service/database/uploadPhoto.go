package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func (db *appdbimpl) UploadPhoto(uploader string, image []byte) (int64, error) {
	id, _ := db.GetUserId(uploader)

	res, err := db.c.Exec("INSERT INTO photos(uploader, uploadTime, likeNumber, commentNumber, image) VALUES(?, CURRENT_TIMESTAMP, 0, 0, ?)", id, image)
	if err != nil {
		panic(err) //TODO
	}
	idPhoto, err := res.LastInsertId()
	if err != nil {
		var photo Photo
		if err := db.c.QueryRow("SELECT photoId FROM photos WHERE photoId=?", idPhoto).Scan(&photo.PhotoId); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return idPhoto, fmt.Errorf("photo does not exist: %w", err)
			}
		}
	}
	return idPhoto, nil
}
