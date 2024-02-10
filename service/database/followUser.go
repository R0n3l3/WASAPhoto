package database

import "log"

func (db *appdbimpl) FollowUser(toFollow string, follower string) (Profile, error) {
	var profile Profile

	idToFollow, err := db.GetUserId(toFollow)
	if err != nil {
		log.Println(err.Error())
		return profile, err
	}

	idFollower, err := db.GetUserId(follower)
	if err != nil {
		log.Println(err.Error())
		return profile, err
	}

	_, err = db.c.Exec("INSERT INTO follow(follower, following) VALUES (?, ?)", idFollower, idToFollow)
	if err != nil {
		log.Println(err.Error())
		return profile, err
	}

	profile, err = db.GetUserProfile(toFollow)
	if err != nil {
		log.Println(err.Error())
		return profile, err
	}
	return profile, nil
}
