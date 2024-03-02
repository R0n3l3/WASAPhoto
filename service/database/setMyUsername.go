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
		log.Println(err.Error())
		return user, err
	}

	if err := db.c.QueryRow("UPDATE users SET username=? WHERE userId=?", new, id).Scan(&user); err != nil {
		log.Println(err.Error())
		return user, err
	}

	if err := db.c.QueryRow("UPDATE profiles SET profileName=? WHERE profileName=?", new, u).Scan(); err != nil {
		if errors.Is(db.c.QueryRow("SELECT profileId FROM profiles WHERE profileName=?", u).Scan(), sql.ErrNoRows) {
			log.Println(err.Error())
			return user, err
		}
		log.Println(err.Error())
		return user, err
	}
	return user, nil
}
