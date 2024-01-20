package database

// AddLike adds a like to the database
func (db *appdbimpl) AddLike(liker, photoID string) error {
    _, err := db.c.Exec("INSERT INTO likes (liker, photo_id) VALUES (?, ?)", liker, photoID)
    return err
}

// DeleteLike removes a like from the database
func (db *appdbimpl) DeleteLike(liker, photoID string) error {
    _, err := db.c.Exec("DELETE FROM likes WHERE liker = ? AND photo_id = ?", liker, photoID)
    return err
}
