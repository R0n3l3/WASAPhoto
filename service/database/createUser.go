package database

import (
	"fmt"
)

func (db *appdbimpl) CreateUser(u string, token uint64) (uint64, error) {
	id, err := db.GetUserId(u)

	if err != nil {
		res, err := db.c.Exec("INSERT INTO profiles(profileName, photoNumber) VALUES (?, 0)", u)
		if err != nil {
			panic(err) //TODO
		}
		idProf, err := res.LastInsertId()
		if err != nil {
			panic(err)
		}
		res, err = db.c.Exec("INSERT INTO users(userId, username, userProfile) VALUES (?, ?, ?)", token, u, idProf)
		if err != nil {
			panic(err)
		}
		idUser, err := res.LastInsertId()
		if err != nil {
			panic(err)
		}
		if idProf != idUser {
			return 0, fmt.Errorf("error in id generation")
		}
		return uint64(idUser), nil
	}
	return id, nil
}
