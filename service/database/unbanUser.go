package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func (db *appdbimpl) UnbanUser(toUnban string, banning string) error {
	idToUnban, _ := db.GetUserId(toUnban)
	idUnbanning, _ := db.GetUserId(banning)

	_, err := db.c.Exec("DELETE FROM ban WHERE banner=? AND banned=?", idUnbanning, idToUnban)
	if err != nil {
		var user1 User
		var user2 User
		if errors.Is(db.c.QueryRow("SELECT banner, banned FROM ban WHERE banner=? AND banned=?", idUnbanning, idToUnban).Scan(&user1.UserId, &user2.UserId), sql.ErrNoRows) {
			return fmt.Errorf("no matching rows found: %w", err)
		}
	}
	return nil
}
