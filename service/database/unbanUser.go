package database

import (
	"database/sql"
	"errors"
	"log"
)

func (db *appdbimpl) UnbanUser(toUnban string, banning string) error {
	idToUnban, err := db.GetUserId(toUnban)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return err
		}
		log.Fatal(err)
	}

	idUnbanning, err := db.GetUserId(banning)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return err
		}
		log.Fatal(err)
	}

	_, err = db.c.Exec("DELETE FROM ban WHERE banner=? AND banned=?", idUnbanning, idToUnban)
	if err != nil {
		var user1 User
		var user2 User
		if errors.Is(db.c.QueryRow("SELECT banner, banned FROM ban WHERE banner=? AND banned=?", idUnbanning, idToUnban).Scan(&user1.UserId, &user2.UserId), sql.ErrNoRows) {
			return err
		}
		log.Fatal(err)
	}
	return nil
}
