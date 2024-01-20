package database

// AddFollow adds a follow relationship to the database
func (db *appdbimpl) AddFollow(follower, followed string) error {
    _, err := db.c.Exec("INSERT INTO follow (follower, followed) VALUES (?, ?)", follower, followed)
    return err
}

// DeleteFollow removes a follow relationship from the database
func (db *appdbimpl) DeleteFollow(follower, followed string) error {
    _, err := db.c.Exec("DELETE FROM follow WHERE follower = ? AND followed = ?", follower, followed)
    return err
}
