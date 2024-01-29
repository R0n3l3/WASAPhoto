package database

func (db *appdbimpl) GetMyStream(u string) ([]int64, error) {
	userId, _ := db.GetUserId(u)
	var photos []int64
	res, err := db.c.Query("SELECT photoId FROM photos p WHERE EXISTS(SELECT * FROM follow WHERE follower=? AND following=p.uploader) ORDER BY uploadTime DESC", userId)
	if err != nil {
		panic(err)
	}
	defer res.Close()
	for res.Next() {
		var id int64
		err = res.Scan(&id)
		if err != nil {
			panic(err)
		}
		photos = append(photos, id)
	}
	return photos, nil
}
