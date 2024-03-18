/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	GetUserId(u string) (uint64, error)
	IsAuthorized(token uint64) bool
	IsBanned(myName string, theirName string) (bool, error)

	DoLogin(u string) (uint64, error)
	GetUserProfile(u string) (Profile, error)
	GetUserProfileId(id uint64) (Profile, error)
	SetMyUsername(u string, new string) error

	BanUser(toBan string, banning string) error
	UnbanUser(toUnban string, unbanning string) error

	GetMyPhotos(name string) ([]Photo, error)
	GetPhoto(id uint64) (Photo, error)
	UploadPhoto(uploader string, image []byte) (Photo, error)
	DeletePhoto(id uint64) error
	GetMyStream(u string) ([]Photo, error)

	GetLike(id uint64) (Like, error)
	GetLikes(photoId uint64) ([]Like, error)
	LikePhoto(photoId uint64, liker string) (Like, error)
	UnlikePhoto(id uint64) error

	GetComment(id uint64) (Comment, error)
	GetComments(id uint64) ([]Comment, error)
	CommentPhoto(photoId uint64, commenter string, content string) (Comment, error)
	UncommentPhoto(commentId uint64) error

	GetFollowers(me string) ([]Profile, error)
	GetFollowing(me string) ([]Profile, error)
	FollowUser(toFollow string, follower string) (Profile, error)
	UnfollowUser(toUnfollow string, unfollower string) error

	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	_, err := db.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		return nil, err
	}
	// Check if table exists. If not, the database is empty, and we need to create the structure

	var tableName string
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='profiles';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		profilesDatabase := `CREATE TABLE profiles (profileId INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, profileName TEXT NOT NULL UNIQUE, photoNumber INTEGER NOT NULL);`
		_, err = db.Exec(profilesDatabase)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='auth';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		authDatabase := `CREATE TABLE auth (token INTEGER NOT NULL PRIMARY KEY);`
		_, err = db.Exec(authDatabase)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='users';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		usersDatabase := `CREATE TABLE users (
    userId INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    userProfile INTEGER NOT NULL UNIQUE,
	FOREIGN KEY (userProfile) REFERENCES profiles(profileId));`
		_, err = db.Exec(usersDatabase)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='ban';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		banDatabase := `CREATE TABLE ban (
    banner INTEGER NOT NULL,
    banned INTEGER NOT NULL,
    PRIMARY KEY (banner, banned),
	FOREIGN KEY (banner) REFERENCES users(userId),
	FOREIGN KEY (banned) REFERENCES users(userId));`
		_, err = db.Exec(banDatabase)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='follow';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		followDatabase := `CREATE TABLE follow (
    follower INTEGER NOT NULL,
    following INTEGER NOT NULL,
    PRIMARY KEY (follower, following),
	FOREIGN KEY (follower) REFERENCES users(userId),
	FOREIGN KEY (following) REFERENCES users(userId));`
		_, err = db.Exec(followDatabase)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='photos';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		photosDatabase := `CREATE TABLE photos (
    photoId INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    uploader INTEGER NOT NULL,
    uploadTime TIMESTAMP NOT NULL,
    likeNumber INTEGER NOT NULL,
    commentNumber INTEGER NOT NULL,
    image BLOB NOT NULL,
	FOREIGN KEY (uploader) REFERENCES users(userId));`
		_, err = db.Exec(photosDatabase)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='likes';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		likesDatabase := `CREATE TABLE likes (
    likeId INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    liker INTEGER NOT NULL,
    photoLiked INTEGER NOT NULL,
	FOREIGN KEY (liker) REFERENCES users(userId),
	FOREIGN KEY(photoLiked) REFERENCES photos(photoId));`
		_, err = db.Exec(likesDatabase)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='comments';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		commentsDatabase := `CREATE TABLE comments (
    commentId INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    commenter TEXT NOT NULL,
    commentTime TIMESTAMP NOT NULL,
    content TEXT NOT NULL,
    photoComment INTEGER NOT NULL,
	FOREIGN KEY (commenter) REFERENCES users(username),
	FOREIGN KEY (photoComment) REFERENCES photos(photoId));`
		_, err = db.Exec(commentsDatabase)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}
	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
