package database

import "errors"

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
func (db *appdbimpl) GetUserDetails(username string) (*UserDetails, error) {
	if username == "" {
		return nil, errors.New("username is empty")
	}
	photosQuery := `SELECT photo_id, image_data FROM photos WHERE username = ? ORDER BY timestamp DESC`
	rows, err := db.c.Query(photosQuery, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var photos []Photo
	for rows.Next() {
		var photo Photo
		if err := rows.Scan(&photo.PhotoID, &photo.ImageData); err != nil {
			return nil, err
		}
		photos = append(photos, photo)
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

	userDetails := &UserDetails{
		Photos:      photos,
		PhotosCount: photosCount,
		Followers:   followersCount,
		Following:   followingCount,
	}

	return userDetails, nil
}
