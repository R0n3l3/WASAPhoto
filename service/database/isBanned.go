package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) IsBanned(myName string, theirName string) bool {
	myId, _ := db.GetUserId(myName)
	theirId, _ := db.GetUserId(theirName)

	if err := db.c.QueryRow("SELECT * FROM ban WHERE banned=? AND banner=?", theirId, myId).Scan(); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			panic(err)
		}
		return false
	}
	return true
}
