package database

func (db *appdbimpl) BanUser(toBan string, banning string) (int64, error) {
	idToBan, _ := db.GetUserId(toBan)
	idBanning, _ := db.GetUserId(banning)

	_, err := db.c.Exec("INSERT INTO ban(banner, banned) VALUES (?, ?)", idBanning, idToBan)
	if err != nil {
		panic(err) //TODO
	}
	return idBanning, nil
}
