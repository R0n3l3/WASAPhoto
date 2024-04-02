package database

import (
	"database/sql"
	"errors"
	"log"
)

func (db *appdbimpl) UnbanUser(toUnban string, banning string) error {
	idToUnban, err := db.GetUserId(toUnban)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Println(err.Error())
		}
		return err
	}

	idUnbanning, err := db.GetUserId(banning)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Println(err.Error())
		}
		return err
	}

	_, err = db.c.Exec("DELETE FROM ban WHERE banner=? AND banned=?", idUnbanning, idToUnban)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Println(err.Error())
		}
		return err
	}
	return nil
}
