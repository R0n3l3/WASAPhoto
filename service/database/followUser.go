package database

import (
	"database/sql"
	"errors"
	"log"
)

func (db *appdbimpl) FollowUser(toFollow string, follower string) (Profile, error) {
	var profile Profile

	idToFollow, err := db.GetUserId(toFollow)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Println(err.Error())
		}
		return profile, err
	}

	idFollower, err := db.GetUserId(follower)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			log.Println(err.Error())
		}
		return profile, err
	}

	var exist int
	err = db.c.QueryRow("SELECT 1 FROM follow WHERE follower=? AND following=?", idFollower, idToFollow).Scan(&exist)
	if errors.Is(err, sql.ErrNoRows) {
		_, err = db.c.Exec("INSERT INTO follow(follower, following) VALUES (?, ?)", idFollower, idToFollow)
		if err != nil {
			log.Println(err.Error())
			return profile, err
		}

		profile, err = db.GetUserProfile(follower)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				log.Println(err.Error())
			}
			return profile, err
		}
		return profile, nil
	} else if err != nil {
		log.Println(err.Error())
		return profile, err
	} else if err == nil {
		return profile, nil
	}
	return profile, err
}
