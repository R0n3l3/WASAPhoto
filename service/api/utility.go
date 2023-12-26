package api

// This file contains some functions that are frequently used by other methods

func getProfile(name string) Profile {
	for i := 0; i < len(Users); i++ {
		if Users[i].Username == name {
			return Users[i].UserProfile
		}
	}
	return Profile{}
}

func getUser(name string) User {
	for i := 0; i < len(Users); i++ {
		if Users[i].Username == name {
			return Users[i]
		}
	}
	return User{}
}
