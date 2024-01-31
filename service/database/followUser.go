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
		if errors.Is(err, sql.ErrNoRows) {
			return profile, err
		}
		log.Fatal(err)
	}

	idFollower, err := db.GetUserId(follower)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return profile, err
		}
		log.Fatal(err)
	}

	_, err = db.c.Exec("INSERT INTO follow(follower, following) VALUES (?, ?)", idFollower, idToFollow)
	if err != nil {
		log.Fatal(err)
	}

	profile, err = db.GetUserProfile(toFollow)
	if err != nil {
		return profile, err
	}
	return profile, nil
}
