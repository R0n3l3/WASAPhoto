package database

import (
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
		log.Println(err.Error())
		return err
	}
	return nil
}
