package database

import (
	"database/sql"
	"errors"
	"log"
	"strings"
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

	_, err = db.c.Exec("INSERT INTO follow(follower, following) VALUES (?, ?)", idFollower, idToFollow)
	if err != nil {
		if !strings.Contains(err.Error(), "UNIQUE constraint failed") {
			log.Println(err.Error())
		}
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
}
