package database

import (
	"database/sql"
	"errors"
	"fmt"
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
    PhotoID   string `json:"photoId" db:"photo_id"`
    ImageData string `json:"imageData" db:"image_data"`
}

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
    GetName() (string, error)
    SetName(name string) error

    Ping() error

    // Example methods for user operations
    GetUser(id string) (*User, error)
    AddUser(user *User) error

    // Add similar methods for Error and Photo
}

type appdbimpl struct {
    c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
func New(db *sql.DB) (AppDatabase, error) {
    // existing implementation...

    // Additional logic for creating tables for User, Error, Photo, etc.
    // Example:
    // _, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (...);`)
    // Handle errors and similar for other entities

    return &appdbimpl{
        c: db,
    }, nil
}

func (db *appdbimpl) Ping() error {
    return db.c.Ping()
}

// GetUser retrieves a user by ID
func (db *appdbimpl) GetUser(id string) (*User, error) {
    // Implement the SQL query to retrieve the user
    // Example: db.c.QueryRow("SELECT ... FROM users WHERE id = ?", id)
    // Scan the result into a User struct and return it
    return nil, errors.New("not implemented")
}

// AddUser adds a new user to the database
func (db *appdbimpl) AddUser(user *User) error {
    // Implement the SQL command to insert a new user
    // Example: _, err := db.c.Exec("INSERT INTO users (...) VALUES (...)", ...)
    // Handle the error and return
    return errors.New("not implemented")
}

// Add similar methods for Error and Photo
