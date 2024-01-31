package database

import (
	"database/sql"
	"errors"
	"log"
)

func (db *appdbimpl) IsBanned(myName string, theirName string) (bool, error) {
	myId, err := db.GetUserId(myName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, err
		}
		log.Fatal(err)
	}

	theirId, err := db.GetUserId(theirName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, err
		}
		log.Fatal(err)
	}

	if err := db.c.QueryRow("SELECT * FROM ban WHERE banned=? AND banner=?", theirId, myId).Scan(); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Fatal(err)
		}
		return false, nil
	}
	return true, nil
}
