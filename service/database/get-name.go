package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) GetName() (string, error) {
	var name string
	err := db.c.QueryRow("SELECT name FROM example_table WHERE id=1").Scan(&name)
	return name, err
}


func (db *appdbimpl) GetUser(userID string) (*User, error) {
	if userID == ""{
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