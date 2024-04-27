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
			statement := "INSERT INTO profiles(profileName, photoNumber) VALUES ('" + u + "', 0)"
			res, err := db.c.Exec(statement)
			if err != nil {
				log.Println(err.Error())
				return id, err
			}

			idProf, err := res.LastInsertId()
			if err != nil {
				log.Println(err.Error())
				return id, err
			}
			prof, _ := db.GetUserProfileId(uint64(idProf))
			res, err = db.c.Exec("INSERT INTO users(username, userProfile) VALUES (?, ?)", prof.ProfileName, uint64(idProf))
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
