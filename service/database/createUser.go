package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) CreateUser(u string, token uint64) (uint64, error) {
	id, err := db.GetUserId(u)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			res, err := db.c.Exec("INSERT INTO profiles(profileName, photoNumber) VALUES (?, 0)", u)
			if err != nil {
				return id, err
			}
			idProf, err := res.LastInsertId()
			if err != nil {
				return id, err
			}
			res, err = db.c.Exec("INSERT INTO users(userId, username, userProfile) VALUES (?, ?, ?)", token, u, idProf)
			if err != nil {
				return id, err
			}
			idUser, err := res.LastInsertId()
			if err != nil {
				return id, err
			}
			return uint64(idUser), nil
		}
		return id, err
	}
	return id, nil
}
