package database

import (
	"log"
)

func (db *appdbimpl) BanUser(toBan string, banning string) (User, error) {
	var user User

	idToBan, err := db.GetUserId(toBan)
	if err != nil {
		log.Println(err.Error())
		return user, err
	}

	idBanning, err := db.GetUserId(banning)
	if err != nil {
		log.Println(err.Error())
		return user, err
	}

	if err := db.c.QueryRow("INSERT INTO ban(banner, banned) VALUES (?, ?)", idBanning, idToBan).Scan(); err != nil {
		log.Println(err.Error())
		return user, err
	}
	if err := db.c.QueryRow("SELECT * FROM users u WHERE EXISTS(SELECT banner FROM ban WHERE banned=? AND banner=? AND u.userId=banner)", idToBan, idBanning).Scan(&user); err != nil {
		log.Println(err.Error())
		return user, err
	}

	return user, nil
}
