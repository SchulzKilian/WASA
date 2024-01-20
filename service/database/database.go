package database

import (
	"database/sql"
    "crypto/rand"
    "encoding/base64"
    "time"
)

// Error represents the error object in the database
type Error struct {
    Error string `json:"error" db:"error"`
}

// User represents the user object in the database
type User struct {
    UserID            string `json:"userId" db:"user_id"`
    Username          string `json:"username" db:"username"`
    Password          string `json:"password" db:"password"`
    Email             string `json:"email" db:"email"`
    Birthday          string `json:"birthday" db:"birthday"`
    SecurityQuestion  string `json:"security_question" db:"security_question"`
    Matricola         int    `json:"matricola" db:"matricola"`
}

// Photo represents the photo object in the database
type Photo struct {
    Username  string `json:"username" db:"username"`
    PhotoID   string `json:"photoId" db:"photo_id"`
    ImageData []byte `json:"imageData" db:"image_data"`
    Timestamp time.Time `json:"timestamp" db:"timestamp"`
}

type Like struct {
    Liker   string `json:"liker" db:"liker"`
    PhotoID string `json:"photoId" db:"photo_id"`
}

type Follow struct {
    Follower   string `json:"follower" db:"follower"`
    Followed string `json:"followed" db:"followed"`
}

type Ban struct {
    Banner   string `json:"banner" db:"banner"`
    Banned string `json:"banned" db:"banned"`
}

type UserDetails struct {
    Photos        []Photo
    PhotosCount   int
    Followers     int
    Following     int
}


// AppDatabase is the high level interface for the DB
type AppDatabase interface {
    GetName() (string, error)
    GetUser(userID string) (*User, error)
    SetName(name string, token string) error
    DoesUserExist(username string) (bool, error, string)
    Ping() error
    GetUserDetails(username string) (*UserDetails, error)

    // Example methods for user operations
    AddUser(user *User) (error, string)

    // Add similar methods for Error and Photo
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
    //Create the table follow
    _, err = db.Exec(`CREATE TABLE IF NOT EXISTS follow (
        follower TEXT,
        followed TEXT
    );`)

    if err != nil{
        return nil, err
    }
    _, err = db.Exec(`CREATE TABLE IF NOT EXISTS likes (
        liker TEXT,
        photo_id TEXT
    );`)

    if err != nil{
        return nil, err
    }
    //create the ban table
    _, err = db.Exec(`CREATE TABLE IF NOT EXISTS bans (
        banner TEXT PRIMARY KEY,
        banned TEXT
    );`)

    if err != nil{
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
        photo_id TEXT PRIMARY KEY,
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