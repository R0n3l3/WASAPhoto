package database

type User struct {
	UserId      uint64 `json:"UserId"`
	Username    string `json:"Username"`
	UserProfile uint64 `json:"UserProfile"`
}

type Profile struct {
	ProfileId   uint64 `json:"ProfileId"`
	ProfileName string `json:"ProfileName"`
	PhotoNumber int    `json:"PhotoNumber"`
}

type Ban struct {
	UserBanned  uint64 `json:"UserBanned"`
	UserBanning uint64 `json:"UserBanning"`
}

type Follow struct {
	Follower  uint64 `json:"Follower"`
	Following uint64 `json:"Following"`
}

type Photo struct {
	PhotoId       uint64 `json:"PhotoId"`
	Uploader      uint64 `json:"Uploader"`
	Image         []byte `json:"Image"`
	UploadTime    string `json:"UploadTime"`
	LikeNumber    int    `json:"LikeNumber"`
	CommentNumber int    `json:"CommentNumber"`
}

type Like struct {
	LikeId     uint64 `json:"LikeId"`
	Liker      uint64 `json:"Liker"`
	PhotoLiked uint64 `json:"PhotoLiked"`
}

type Comment struct {
	CommentId    uint64 `json:"CommentId"`
	Commenter    string `json:"Commenter"`
	CommentTime  string `json:"CommentTime"`
	Content      string `json:"Content"`
	PhotoComment uint64 `json:"PhotoComment"`
}
