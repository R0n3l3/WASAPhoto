package database

import (
	"database/sql"
	"errors"
	"log"
)

func (db *appdbimpl) DoLogin(u string) ([]Profile, error) {
	var users []Profile
	users, err := db.GetUserProfiles(u)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			statement := `INSERT INTO profiles(profileName, photoNumber) VALUES ("` + u + `", 0)`
			res, err := db.c.Exec(statement)
			if err != nil {
				log.Println(err.Error())
				return users, err
			}
			idProf, err := res.LastInsertId()
			if err != nil {
				log.Println(err.Error())
				return users, err
			}
			prof, _ := db.GetUserProfileId(uint64(idProf))
			res, err = db.c.Exec("INSERT INTO users(username, userProfile) VALUES (?, ?)", prof.ProfileName, uint64(idProf))
			if err != nil {
				log.Println(err.Error())
				return users, err
			}
			idUser, err := res.LastInsertId()
			if err != nil {
				log.Println(err.Error())
				return users, err
			}
			res, err = db.c.Exec("INSERT INTO auth(token) VALUES (?)", uint64(idUser))
			if err != nil {
				log.Println(err.Error())
				return users, err
			}
			users = append(users, prof)
			return users, nil
		}
		log.Println(err.Error())
		return users, err
	}
	return users, nil
}
