package database

import (
	"database/sql"
	"errors"
	"log"
	"strings"
)

func (db *appdbimpl) BanUser(toBan string, banning string) error {

	idToBan, err := db.GetUserId(toBan)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Println(err.Error())
		}
		return err
	}

	idBanning, err := db.GetUserId(banning)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Println(err.Error())
		}
		return err
	}

	_, err = db.c.Exec("INSERT INTO ban(banner, banned) VALUES (?, ?)", idBanning, idToBan)
	if err != nil {
		if !strings.Contains(err.Error(), "UNIQUE constraint failed") {
			log.Println(err.Error())
		}
		return err
	}

	return nil
}
