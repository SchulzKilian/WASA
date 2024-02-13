package database

import (
	"database/sql"
	"encoding/base64"
	"errors"
)

// GetName is an example that shows you how to query data
func (db *appdbimpl) GetName() (string, error) {
	var name string
	err := db.c.QueryRow("SELECT name FROM example_table WHERE id=1").Scan(&name)
	return name, err
}

func (db *appdbimpl) GetUser(userID string) (*User, error) {
	if userID == "" {
		return nil, nil
	}
	// Prepare a user instance to hold the data
	var user User

	// SQL query to select the user by user_id
	query := "SELECT user_id, username, password, email, birthday, security_question, matricola FROM users WHERE user_id = ?"

	// Execute the query
	err := db.c.QueryRow(query, userID).Scan(&user.UserID, &user.Username, &user.Password, &user.Email, &user.Birthday, &user.SecurityQuestion, &user.Matricola)
	if err != nil {
		// Other error occurred
		return nil, err
	}

	// Return the user object if found
	return &user, nil
}
func (db *appdbimpl) GetUserDetails(username, currentUsername string) (*UserDetails, error) {
	if username == "" {
		return nil, errors.New("username is empty")
	}
	query := "SELECT username FROM users WHERE username = ?"
	var user User
	err := db.c.QueryRow(query, username).Scan(&user.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {

			return nil, nil
		}
		// Other error occurred, return the error
		return nil, err
	}
	photosQuery := `SELECT photo_id, image_data FROM photos WHERE username = ? ORDER BY timestamp DESC`
	rows, err := db.c.Query(photosQuery, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	userDetails := &UserDetails{}
	for rows.Next() {
		var photo PhotoDetails
		var imageData []byte
		if err := rows.Scan(&photo.PhotoID, &imageData); err != nil {
			return nil, err
		}
		photo.Username = username
		photo.ImageData = base64.StdEncoding.EncodeToString(imageData)

		// Query to count likes for the current photo
		var likesCount int
		err := db.c.QueryRow("SELECT COUNT(*) FROM likes WHERE photo_id = ?", photo.PhotoID).Scan(&likesCount)
		if err != nil {
			return nil, err // Handle error appropriately
		}
		photo.LikesCount = likesCount

		var Liking bool
		err = db.c.QueryRow(`SELECT EXISTS (
			SELECT 1
			FROM likes
			WHERE liker = ?
			AND photo_id = ?
		)`, currentUsername, photo.PhotoID).Scan(&Liking)
		if err != nil {
			return nil, err // Handle error appropriately
		}
		photo.Liking = Liking

		// Query to count comments for the current photo
		var commentsCount int
		err = db.c.QueryRow("SELECT COUNT(*) FROM comments WHERE photo_id = ?", photo.PhotoID).Scan(&commentsCount)
		if err != nil {
			return nil, err // Handle error appropriately
		}
		photo.CommentsCount = commentsCount

		userDetails.Photos = append(userDetails.Photos, photo)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Get count of photos
	photosCountQuery := `SELECT COUNT(*) FROM photos WHERE username = ?`
	var photosCount int
	err = db.c.QueryRow(photosCountQuery, username).Scan(&photosCount)
	if err != nil {
		return nil, err
	}

	// Get count of followers
	followersQuery := `SELECT COUNT(*) FROM follow WHERE followed = ?`
	var followersCount int
	err = db.c.QueryRow(followersQuery, username).Scan(&followersCount)
	if err != nil {
		return nil, err
	}

	// Get count of following
	followingQuery := `SELECT COUNT(*) FROM follow WHERE follower = ?`
	var followingCount int
	err = db.c.QueryRow(followingQuery, username).Scan(&followingCount)
	if err != nil {
		return nil, err
	}
	var isBanning bool
	isBanningQuery := `SELECT EXISTS (SELECT 1 FROM bans WHERE banner = ? AND banned = ?)`
	err = db.c.QueryRow(isBanningQuery, currentUsername, username).Scan(&isBanning)
	if err != nil {
		return nil, err
	}
	isFollowingQuery := `SELECT EXISTS (SELECT 1 FROM follow WHERE follower = ? AND followed = ?)`
	var isFollowing bool
	err = db.c.QueryRow(isFollowingQuery, currentUsername, username).Scan(&isFollowing)
	if err != nil {
		return nil, err
	}

	userDetails.IsFollowing = isFollowing
	userDetails.Followers = followersCount
	userDetails.Following = followingCount
	userDetails.IsBanning = isBanning

	return userDetails, nil
}
