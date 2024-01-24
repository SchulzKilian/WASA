package database

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"time"
)

// Error represents the error object in the database
type Error struct {
	Error string `json:"error" db:"error"`
}

// User represents the user object in the database
type User struct {
	UserID           string `json:"userId" db:"user_id"`
	Username         string `json:"username" db:"username"`
	Password         string `json:"password" db:"password"`
	Email            string `json:"email" db:"email"`
	Birthday         string `json:"birthday" db:"birthday"`
	SecurityQuestion string `json:"security_question" db:"security_question"`
	Matricola        int    `json:"matricola" db:"matricola"`
}

// Photo represents the photo object in the database
type Photo struct {
	Username  string    `json:"username" db:"username"`
	PhotoID   string    `json:"photoId" db:"photo_id"`
	ImageData []byte    `json:"imageData" db:"image_data"`
	Timestamp time.Time `json:"timestamp" db:"timestamp"`
}

type PhotoDetails struct {
	Username      string    `json:"username"`
	PhotoID       string    `json:"photoId"`
	ImageData     []byte    `json:"imageData"`
	Timestamp     time.Time `json:"timestamp"`
	LikesCount    int       `json:"LikesCount"`
	CommentsCount int       `json:"CommentsCount"`
}

type Like struct {
	Liker   string `json:"liker" db:"liker"`
	PhotoID string `json:"photoId" db:"photo_id"`
}

type Comment struct {
	Content   string `json:"content" db:"content"`
	Commenter string `json:"commenter" db:"commenter"`
	PhotoID   string `json:"photoId" db:"photo_id"`
}

type Follow struct {
	Follower string `json:"follower" db:"follower"`
	Followed string `json:"followed" db:"followed"`
}

type Ban struct {
	Banner string `json:"banner" db:"banner"`
	Banned string `json:"banned" db:"banned"`
}

type UserDetails struct {
	Photos      []PhotoDetails
	PhotosCount int
	Followers   int
	Following   int
	IsFollowing bool `json:"IsFollowing"`
}

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	GetName() (string, error)
	GetUser(userID string) (*User, error)
	SetName(name string, token string) error
	DoesUserExist(username string) (bool, error, string)
	AddFollow(follower, followed string) error
	DeleteFollow(follower, followed string) error
	Ping() error
	GetUserDetails(username, currentusername string) (*UserDetails, error)
	AddPhoto(photo Photo) error
	DeletePhoto(photo_id int) error
	AddComment(comment Comment) error
	DeleteComment(commentID, commenter string) error
	AddLike(liker, photoID string) error
	DeleteLike(liker, photoID string) error
	AddBan(banner, banned string) error
	DeleteBan(banner, banned string) error
	AddUser(user *User) (error, string)
	GetFollowedUsersPhotos(username string) ([]PhotoDetails, error)
	AmIBanned(banner, banned string) (bool, error)
	Close() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
func New(db *sql.DB) (AppDatabase, error) {
	// Create Error table
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS errors (
        error TEXT
    );`)
	if err != nil {
		return nil, err
	}
	// Create the table follow
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS follow (
        follower TEXT,
        followed TEXT,
        PRIMARY KEY (follower, followed)
    );`)

	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS likes (
        liker TEXT,
        photo_id TEXT,
        PRIMARY KEY (liker, photo_id)
    );`)

	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS comments (
        comment_id INTEGER PRIMARY KEY AUTOINCREMENT,
        content TEXT,
        commenter TEXT,
        photo_id TEXT
    );`)

	if err != nil {
		return nil, err
	}
	// create the ban table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS bans (
        banner TEXT PRIMARY KEY,
        banned TEXT
    );`)

	if err != nil {
		return nil, err
	}

	// Create User table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
        user_id TEXT UNIQUE,
        username TEXT PRIMARY KEY,
        password TEXT,
        email TEXT,
        birthday TEXT,
        security_question TEXT,
        matricola INTEGER
    );`)
	if err != nil {
		return nil, err
	}

	// Create Photo table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS photos (
        username TEXT,
        photo_id INTEGER PRIMARY KEY AUTOINCREMENT,
        image_data BLOB,
        timestamp TIMESTAMP
    );`)
	if err != nil {
		return nil, err
	}

	// Continue with any additional setup
	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
func (db *appdbimpl) Close() error {
	if db.c != nil {
		return db.c.Close()
	}
	return nil
}

// GetUser retrieves a user by ID

// AddUser adds a new user to the database
func (db *appdbimpl) AddUser(user *User) (error, string) {
	var err error
	user.UserID, err = generateRandomString(10)
	if err != nil {
		return err, ""
	}
	_, err = db.c.Exec("INSERT INTO users (user_id, username, password, email, birthday, security_question, matricola) VALUES (?, ?, ?, ?, ?, ?, ?)",
		user.UserID, user.Username, user.Password, user.Email, user.Birthday, user.SecurityQuestion, user.Matricola)
	return err, user.UserID
}

func (db *appdbimpl) DoesUserExist(username string) (bool, error, string) {
	var userID string
	query := "SELECT user_id FROM users WHERE username = ? LIMIT 1"

	err := db.c.QueryRow(query, username).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil, ""
		}
		return false, err, ""
	}

	return true, nil, userID
}

func generateRandomString(n int) (string, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b)[:n], nil
}

func (db *appdbimpl) GetFollowedUsersPhotos(username string) ([]PhotoDetails, error) {
	var photos []PhotoDetails

	// SQL query to retrieve photos, along with likes and comments count
	query := `SELECT p.username, p.photo_id, p.image_data, p.timestamp, 
                     COUNT(DISTINCT l.liker) AS likes_count, 
                     COUNT(DISTINCT c.comment_id) AS comments_count 
              FROM photos p
              LEFT JOIN likes l ON p.photo_id = l.photo_id
              LEFT JOIN comments c ON p.photo_id = c.photo_id
              JOIN follow f ON p.username = f.followed
              WHERE f.follower = ?
              GROUP BY p.photo_id`

	rows, err := db.c.Query(query, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var photo PhotoDetails
		err := rows.Scan(&photo.Username, &photo.PhotoID, &photo.ImageData, &photo.Timestamp, &photo.LikesCount, &photo.CommentsCount)
		if err != nil {
			return nil, err
		}
		photos = append(photos, photo)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return photos, nil
}
