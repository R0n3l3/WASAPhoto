package database

type User struct {
	UserId      int64
	Username    string
	UserProfile int64
}

type Profile struct {
	ProfileId   int64
	ProfileName string
	PhotoNumber int
}

type Ban struct {
	UserBanned  int64
	UserBanning int64
}

type Follow struct {
	Follower  int64
	Following int64
}

type Photo struct { // Create a new photo
	PhotoId       int64
	Uploader      int64
	Image         []byte
	UploadTime    string
	LikeNumber    int
	CommentNumber int
}

type Like struct {
	LikeId     int64
	Liker      int64
	PhotoLiked int64
}

type Comment struct {
	CommentId    int64
	Commenter    int64
	CommentTime  string
	Content      string
	PhotoComment int64
}
