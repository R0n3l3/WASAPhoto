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
	CreateUser(u string) (int64, error)
	SetMyUsername(u string, new string) (User, error)
	GetUserId(u string) (int64, error)
	GetUserProfile(u string) (Profile, error)

	BanUser(toBan string, banning string) (User, error)
	UnbanUser(toUnban string, unbanning string) error
	IsBanned(myName string, theirName string) bool

	UploadPhoto(uploader string, image []byte) (Photo, error)
	DeletePhoto(id int64) error
	GetMyStream(u string) ([]int64, error)

	LikePhoto(photoId int64, liker string) (Like, error)
	UnlikePhoto(id int64) error

	CommentPhoto(photoId int64, commenter string, content string) (int64, error)
	UncommentPhoto(commentId int64) error

	FollowUser(toFollow string, follower string) (int64, error)
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
    commenter INTEGER NOT NULL, 
    commentTime TIMESTAMP NOT NULL, 
    content TEXT NOT NULL, 
    photoComment INTEGER NOT NULL, 
	FOREIGN KEY (commenter) REFERENCES users(userId), 
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
