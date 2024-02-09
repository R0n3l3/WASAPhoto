package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) UnbanUser(toUnban string, banning string) error {
	idToUnban, err := db.GetUserId(toUnban)
	if err != nil {
		print(err.Error())
		return err
	}

	idUnbanning, err := db.GetUserId(banning)
	if err != nil {
		print(err.Error())
		return err
	}

	_, err = db.c.Exec("DELETE FROM ban WHERE banner=? AND banned=?", idUnbanning, idToUnban)
	if err != nil {
		var user1 User
		var user2 User
		if errors.Is(db.c.QueryRow("SELECT banner, banned FROM ban WHERE banner=? AND banned=?", idUnbanning, idToUnban).Scan(&user1.UserId, &user2.UserId), sql.ErrNoRows) {
			print(err.Error())
			return err
		}
		print(err.Error())
		return err
	}
	return nil
}
