package api

type Like struct {
	LikeId int
	Liker  User
}

type Comment struct {
	CommentId   int
	Commenter   User
	Content     string
	CommentTime string
}

type Image struct {
	Link string
}

type Photo struct {
	PhotoId       int
	Uploader      User
	PhotoImage    Image
	UploadTime    string
	Likes         []Like
	LikeNumber    int
	Comments      []Comment
	CommentNumber int
}

type Profile struct {
	ProfileId   string
	Photos      []Photo
	PhotoNumber int
	Followers   []User
	Following   []User
}

type User struct {
	Username    string
	Banned      []User
	UserProfile Profile
}
