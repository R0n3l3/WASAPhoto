package database

import (
	"database/sql"
	"errors"
	"log"
)

func (db *appdbimpl) IsBanned(myName string, theirName string) (bool, error) {
	myId, err := db.GetUserId(myName)
	if err != nil {
		log.Println(err.Error())
		return false, err
	}

	theirId, err := db.GetUserId(theirName)
	if err != nil {
		log.Println(err.Error())
		return false, err
	}
	if err := db.c.QueryRow("SELECT * FROM ban WHERE banned=? AND banner=?", theirId, myId).Scan(); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Println(err.Error())
			return false, err
		}
		return false, nil
	}
	return true, nil
}
