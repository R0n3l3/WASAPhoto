package database

import (
	"database/sql"
	"errors"
	"log"
)

func (db *appdbimpl) BanUser(toBan string, banning string) (User, error) {
	var user User

	idToBan, err := db.GetUserId(toBan)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, err
		}
		log.Fatal(err)
	}

	idBanning, err := db.GetUserId(banning)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, err
		}
		log.Fatal(err)
	}

	if err := db.c.QueryRow("INSERT INTO ban(banner, banned) VALUES (?, ?)", idBanning, idToBan).Scan(); err != nil {
		log.Fatal(err)
	}
	if err := db.c.QueryRow("SELECT * FROM users u WHERE EXISTS(SELECT banner FROM ban WHERE banned=? AND banner=? AND u.userId=banner)", idToBan, idBanning).Scan(&user); err != nil {
		log.Fatal(err)
	}

	return user, nil
}
