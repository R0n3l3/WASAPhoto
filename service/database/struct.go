package database

type User struct {
	UserId      uint64
	Username    string
	UserProfile uint64
}

type Profile struct {
	ProfileId   uint64
	ProfileName string
	PhotoNumber int
}

type Ban struct {
	UserBanned  uint64
	UserBanning uint64
}

type Follow struct {
	Follower  uint64
	Following uint64
}

type Photo struct { // Create a new photo
	PhotoId       uint64
	Uploader      uint64
	Image         []byte
	UploadTime    string
	LikeNumber    int
	CommentNumber int
}

type Like struct {
	LikeId     uint64
	Liker      uint64
	PhotoLiked uint64
}

type Comment struct {
	CommentId    uint64
	Commenter    uint64
	CommentTime  string
	Content      string
	PhotoComment uint64
}
