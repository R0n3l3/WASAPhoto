package database

import (
	"database/sql"
	"errors"
	"log"
)

func (db *appdbimpl) DoLogin(u string) (uint64, error) {
	id, err := db.GetUserId(u)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			res, err := db.c.Exec("INSERT INTO profiles(profileName, photoNumber) VALUES (?, 0)", u)
			if err != nil {
				log.Println(err.Error())
				return id, err
			}
			idProf, err := res.LastInsertId()
			if err != nil {
				log.Println(err.Error())
				return id, err
			}
			res, err = db.c.Exec("INSERT INTO users(username, userProfile) VALUES (?, ?)", u, idProf)
			if err != nil {
				log.Println(err.Error())
				return id, err
			}
			idUser, err := res.LastInsertId()
			if err != nil {
				log.Println(err.Error())
				return id, err
			}
			res, err = db.c.Exec("INSERT INTO auth(token) VALUES (?)", uint64(idUser))
			if err != nil {
				log.Println(err.Error())
				return id, err
			}
			return uint64(idUser), nil
		}
		log.Println(err.Error())
		return id, err
	}
	return id, nil
}
