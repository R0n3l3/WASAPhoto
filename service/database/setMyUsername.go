package database

import (
	"database/sql"
	"errors"
	"log"
)

func (db *appdbimpl) SetMyUsername(u string, new string) (User, error) {
	var user User

	id, err := db.GetUserId(u)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, err
		}
		log.Fatal(err)
	}

	if err := db.c.QueryRow("UPDATE users SET username=? WHERE userId=?", new, id).Scan(&user); err != nil {
		log.Fatal(err)
	}

	if err := db.c.QueryRow("UPDATE profiles SET profileName=? WHERE profileName=?", new, u).Scan(); err != nil {
		if errors.Is(db.c.QueryRow("SELECT profileId FROM profiles WHERE profileName=?", u).Scan(), sql.ErrNoRows) {
			return user, err
		}
		log.Fatal(err)
	}
	return user, nil
}
