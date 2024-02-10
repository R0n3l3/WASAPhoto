package database

import (
	"database/sql"
	"errors"
	"log"
)

func (db *appdbimpl) UnfollowUser(toUnfollow string, unfollower string) error {
	idToUnfollow, err := db.GetUserId(toUnfollow)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	idUnfollower, err := db.GetUserId(unfollower)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	_, err = db.c.Exec("DELETE FROM follow WHERE follower=? AND following=?", idUnfollower, idToUnfollow)
	if err != nil {
		var user1 User
		var user2 User
		if errors.Is(db.c.QueryRow("SELECT follower, following FROM follow WHERE follower=? AND following=?", idUnfollower, idToUnfollow).Scan(&user1.UserId, &user2.UserId), sql.ErrNoRows) {
			log.Println(err.Error())
			return err
		}
		log.Println(err.Error())
		return err
	}
	return nil
}
