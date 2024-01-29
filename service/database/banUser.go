package database

func (db *appdbimpl) BanUser(toBan string, banning string) (User, error) {
	idToBan, _ := db.GetUserId(toBan)
	idBanning, _ := db.GetUserId(banning)

	var user User

	if err := db.c.QueryRow("INSERT INTO ban(banner, banned) VALUES (?, ?)", idBanning, idToBan).Scan(&user); err != nil {
		panic(err) //TODO
	}
	return user, nil
}
