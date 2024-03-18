package database

import (
	"log"
)

func (db *appdbimpl) BanUser(toBan string, banning string) error {

	idToBan, err := db.GetUserId(toBan)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	idBanning, err := db.GetUserId(banning)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	_, err = db.c.Exec("INSERT INTO ban(banner, banned) VALUES (?, ?)", idBanning, idToBan)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
