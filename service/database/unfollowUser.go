package database

import (
	"database/sql"
	"errors"
	"fmt"
)

func (db *appdbimpl) UnfollowUser(toUnfollow string, unfollower string) error {
	idToUnfollow, _ := db.GetUserId(toUnfollow)
	idUnfollower, _ := db.GetUserId(unfollower)

	_, err := db.c.Exec("DELETE FROM follow WHERE follower=? AND following=?", idUnfollower, idToUnfollow)
	if err != nil {
		var user1 User
		var user2 User
		if errors.Is(db.c.QueryRow("SELECT follower, following FROM follow WHERE follower=? AND following=?", idUnfollower, idToUnfollow).Scan(&user1.UserId, &user2.UserId), sql.ErrNoRows) {
			return fmt.Errorf("no matching rows found: %w", err)
		}
	}
	return nil
}
