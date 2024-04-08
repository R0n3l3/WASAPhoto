package database

import (
	"database/sql"
	"errors"
	"log"
	"strings"
)

func (db *appdbimpl) SetMyUsername(u string, new string) error {
	id, err := db.GetUserId(u)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Println(err.Error())
		}
		return err
	}

	_, err = db.c.Exec("UPDATE profiles SET profileName=? WHERE profileName=?", new, u)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Println(err.Error())
		}
		return err
	}

	_, err = db.c.Exec("UPDATE comments SET commenter=? WHERE commenter=?", new, u)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Println(err.Error())
			return err
		}
	}

	_, err = db.c.Exec("UPDATE users SET username=? WHERE userId=?", new, id)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) && !strings.Contains(err.Error(), "UNIQUE constraint failed") {
			log.Println(err.Error())
		}
		return err
	}
	return nil
}
